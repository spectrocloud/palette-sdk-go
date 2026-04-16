// Package awsinfra provides direct AWS infrastructure queries using the AWS SDK.
// Credentials are sourced from Palette cloud accounts, not env vars.
package awsinfra

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	elbtypes "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

// Service defines operations for querying and managing AWS infrastructure.
type Service interface {
	// Network
	ListVPCs(creds Credentials) ([]VPC, error)
	ListSubnets(creds Credentials, vpcID string) ([]Subnet, error)
	ListKeyPairs(creds Credentials) ([]KeyPair, error)
	ListSecurityGroups(creds Credentials, vpcID string) ([]SecurityGroup, error)
	ListAvailabilityZones(creds Credentials) ([]string, error)
	ListInstanceTypes(creds Credentials) ([]string, error)

	// EKS
	ListEKSClusters(creds Credentials) ([]string, error)

	// EC2 Instances
	ListInstances(creds Credentials, filters map[string]string) ([]Instance, error)
	GetInstance(creds Credentials, instanceID string) (*Instance, error)
	LaunchInstance(creds Credentials, spec LaunchSpec) (*Instance, error)
	StartInstance(creds Credentials, instanceID string) error
	StopInstance(creds Credentials, instanceID string) error
	RebootInstance(creds Credentials, instanceID string) error
	TerminateInstance(creds Credentials, instanceID string) error

	// Security Groups
	CreateSecurityGroup(creds Credentials, spec SGSpec) (*SecurityGroup, error)

	// Load Balancer
	CreateNLB(creds Credentials, spec NLBSpec) (*NLBResult, error)
	DeleteNLB(creds Credentials, arn string) error
}

// SGSpec defines parameters for creating a security group.
type SGSpec struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	VPCID       string   `json:"vpcId"`
	Rules       []SGRule `json:"rules"`
}

// NLBSpec defines parameters for creating a Network Load Balancer with target group.
type NLBSpec struct {
	Name        string   `json:"name"`
	VPCID       string   `json:"vpcId"`
	SubnetIDs   []string `json:"subnetIds"`  // Subnets for the NLB
	Port        int32    `json:"port"`        // Listener + target port (e.g., 6443)
	TargetIDs   []string `json:"targetIds"`   // EC2 instance IDs to register
}

// NLBResult contains the created NLB details.
type NLBResult struct {
	LBARN        string `json:"lbArn"`
	LBName       string `json:"lbName"`
	LBDNS        string `json:"lbDns"`
	TGARN        string `json:"tgArn"`
	TGName       string `json:"tgName"`
	Port         int32  `json:"port"`
	TargetCount  int    `json:"targetCount"`
}

// SGRule defines an ingress or egress rule.
type SGRule struct {
	Direction string `json:"direction"` // "ingress" or "egress"
	Protocol  string `json:"protocol"`  // "tcp", "udp", "-1" (all)
	FromPort  int32  `json:"fromPort"`
	ToPort    int32  `json:"toPort"`
	CIDR      string `json:"cidr"` // e.g., "0.0.0.0/0" or "10.0.0.0/8"
}

// LaunchSpec defines parameters for launching a new EC2 instance.
type LaunchSpec struct {
	Name           string `json:"name"`
	ImageID        string   `json:"imageId"`        // AMI ID
	InstanceType   string   `json:"instanceType"`   // e.g., t3.large
	KeyName        string   `json:"keyName"`        // SSH key pair name
	SubnetID       string   `json:"subnetId"`       // Subnet to launch in
	SecurityGroups []string `json:"securityGroups"` // Security group IDs
	DiskSizeGB     int32    `json:"diskSizeGb"`     // Root volume size
	UserData       string   `json:"userData"`       // Base64-encoded or raw user-data script
}

// Credentials sourced from a Palette cloud account.
type Credentials struct {
	AccessKey string
	SecretKey string
	Region    string
}

// VPC represents an AWS VPC.
type VPC struct {
	ID        string `json:"id"`
	CIDR      string `json:"cidr"`
	Name      string `json:"name"`
	IsDefault bool   `json:"isDefault"`
	State     string `json:"state"`
}

// Subnet represents an AWS subnet.
type Subnet struct {
	ID               string `json:"id"`
	CIDR             string `json:"cidr"`
	Name             string `json:"name"`
	AvailabilityZone string `json:"availabilityZone"`
	Public           bool   `json:"public"`
	VPCID            string `json:"vpcId"`
}

// KeyPair represents an AWS key pair.
type KeyPair struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// SecurityGroup represents an AWS security group.
type SecurityGroup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VPCID       string `json:"vpcId"`
}

// Instance represents an EC2 instance.
type Instance struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	State            string `json:"state"`
	Type             string `json:"type"`
	PrivateIP        string `json:"privateIp"`
	PublicIP         string `json:"publicIp"`
	VPCID            string `json:"vpcId"`
	SubnetID         string `json:"subnetId"`
	AvailabilityZone string `json:"availabilityZone"`
	LaunchTime       string `json:"launchTime"`
	KeyName          string `json:"keyName"`
}

type service struct{}

// New creates a new AWS infra service.
func New() Service {
	return &service{}
}

func (s *service) newEC2Client(creds Credentials) *ec2.Client {
	return ec2.New(ec2.Options{
		Region:      creds.Region,
		Credentials: credentials.NewStaticCredentialsProvider(creds.AccessKey, creds.SecretKey, ""),
	})
}

func (s *service) newELBClient(creds Credentials) *elbv2.Client {
	return elbv2.New(elbv2.Options{
		Region:      creds.Region,
		Credentials: credentials.NewStaticCredentialsProvider(creds.AccessKey, creds.SecretKey, ""),
	})
}

func (s *service) newEKSClient(creds Credentials) *eks.Client {
	return eks.New(eks.Options{
		Region:      creds.Region,
		Credentials: credentials.NewStaticCredentialsProvider(creds.AccessKey, creds.SecretKey, ""),
	})
}

func nameFromTags(tags []types.Tag) string {
	for _, t := range tags {
		if aws.ToString(t.Key) == "Name" {
			return aws.ToString(t.Value)
		}
	}
	return ""
}

func (s *service) ListVPCs(creds Credentials) ([]VPC, error) {
	client := s.newEC2Client(creds)
	out, err := client.DescribeVpcs(context.Background(), &ec2.DescribeVpcsInput{})
	if err != nil {
		return nil, err
	}
	vpcs := make([]VPC, 0, len(out.Vpcs))
	for _, v := range out.Vpcs {
		vpcs = append(vpcs, VPC{
			ID:        aws.ToString(v.VpcId),
			CIDR:      aws.ToString(v.CidrBlock),
			Name:      nameFromTags(v.Tags),
			IsDefault: aws.ToBool(v.IsDefault),
			State:     string(v.State),
		})
	}
	return vpcs, nil
}

func (s *service) ListSubnets(creds Credentials, vpcID string) ([]Subnet, error) {
	client := s.newEC2Client(creds)
	input := &ec2.DescribeSubnetsInput{}
	if vpcID != "" {
		input.Filters = []types.Filter{
			{Name: aws.String("vpc-id"), Values: []string{vpcID}},
		}
	}
	out, err := client.DescribeSubnets(context.Background(), input)
	if err != nil {
		return nil, err
	}
	subnets := make([]Subnet, 0, len(out.Subnets))
	for _, sub := range out.Subnets {
		subnets = append(subnets, Subnet{
			ID:               aws.ToString(sub.SubnetId),
			CIDR:             aws.ToString(sub.CidrBlock),
			Name:             nameFromTags(sub.Tags),
			AvailabilityZone: aws.ToString(sub.AvailabilityZone),
			Public:           aws.ToBool(sub.MapPublicIpOnLaunch),
			VPCID:            aws.ToString(sub.VpcId),
		})
	}
	return subnets, nil
}

func (s *service) ListKeyPairs(creds Credentials) ([]KeyPair, error) {
	client := s.newEC2Client(creds)
	out, err := client.DescribeKeyPairs(context.Background(), &ec2.DescribeKeyPairsInput{})
	if err != nil {
		return nil, err
	}
	pairs := make([]KeyPair, 0, len(out.KeyPairs))
	for _, kp := range out.KeyPairs {
		pairs = append(pairs, KeyPair{
			Name: aws.ToString(kp.KeyName),
			Type: string(kp.KeyType),
		})
	}
	return pairs, nil
}

func (s *service) ListSecurityGroups(creds Credentials, vpcID string) ([]SecurityGroup, error) {
	client := s.newEC2Client(creds)
	input := &ec2.DescribeSecurityGroupsInput{}
	if vpcID != "" {
		input.Filters = []types.Filter{
			{Name: aws.String("vpc-id"), Values: []string{vpcID}},
		}
	}
	out, err := client.DescribeSecurityGroups(context.Background(), input)
	if err != nil {
		return nil, err
	}
	groups := make([]SecurityGroup, 0, len(out.SecurityGroups))
	for _, sg := range out.SecurityGroups {
		groups = append(groups, SecurityGroup{
			ID:          aws.ToString(sg.GroupId),
			Name:        aws.ToString(sg.GroupName),
			Description: aws.ToString(sg.Description),
			VPCID:       aws.ToString(sg.VpcId),
		})
	}
	return groups, nil
}

func (s *service) ListInstanceTypes(creds Credentials) ([]string, error) {
	// Return common instance types — a full list is huge
	return []string{
		"t3.micro", "t3.small", "t3.medium", "t3.large", "t3.xlarge", "t3.2xlarge",
		"m5.large", "m5.xlarge", "m5.2xlarge", "m5.4xlarge",
		"c5.large", "c5.xlarge", "c5.2xlarge", "c5.4xlarge",
		"r5.large", "r5.xlarge", "r5.2xlarge",
		"m6i.large", "m6i.xlarge", "m6i.2xlarge",
		"c6i.large", "c6i.xlarge", "c6i.2xlarge",
	}, nil
}

func (s *service) ListAvailabilityZones(creds Credentials) ([]string, error) {
	client := s.newEC2Client(creds)
	out, err := client.DescribeAvailabilityZones(context.Background(), &ec2.DescribeAvailabilityZonesInput{})
	if err != nil {
		return nil, err
	}
	zones := make([]string, 0, len(out.AvailabilityZones))
	for _, az := range out.AvailabilityZones {
		zones = append(zones, aws.ToString(az.ZoneName))
	}
	return zones, nil
}

func (s *service) ListEKSClusters(creds Credentials) ([]string, error) {
	client := s.newEKSClient(creds)
	out, err := client.ListClusters(context.Background(), &eks.ListClustersInput{})
	if err != nil {
		return nil, err
	}
	return out.Clusters, nil
}

func (s *service) ListInstances(creds Credentials, filters map[string]string) ([]Instance, error) {
	client := s.newEC2Client(creds)
	input := &ec2.DescribeInstancesInput{}
	if len(filters) > 0 {
		var f []types.Filter
		for k, v := range filters {
			f = append(f, types.Filter{Name: aws.String(k), Values: []string{v}})
		}
		input.Filters = f
	}
	out, err := client.DescribeInstances(context.Background(), input)
	if err != nil {
		return nil, err
	}
	var instances []Instance
	for _, r := range out.Reservations {
		for _, i := range r.Instances {
			instances = append(instances, Instance{
				ID:               aws.ToString(i.InstanceId),
				Name:             nameFromTags(i.Tags),
				State:            string(i.State.Name),
				Type:             string(i.InstanceType),
				PrivateIP:        aws.ToString(i.PrivateIpAddress),
				PublicIP:         aws.ToString(i.PublicIpAddress),
				VPCID:            aws.ToString(i.VpcId),
				SubnetID:         aws.ToString(i.SubnetId),
				AvailabilityZone: aws.ToString(i.Placement.AvailabilityZone),
				LaunchTime:       i.LaunchTime.Format("2006-01-02 15:04"),
				KeyName:          aws.ToString(i.KeyName),
			})
		}
	}
	return instances, nil
}

func (s *service) LaunchInstance(creds Credentials, spec LaunchSpec) (*Instance, error) {
	client := s.newEC2Client(creds)

	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(spec.ImageID),
		InstanceType: types.InstanceType(spec.InstanceType),
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	}

	if spec.KeyName != "" {
		input.KeyName = aws.String(spec.KeyName)
	}
	if spec.SubnetID != "" {
		input.SubnetId = aws.String(spec.SubnetID)
	}
	if len(spec.SecurityGroups) > 0 {
		input.SecurityGroupIds = spec.SecurityGroups
	}
	if spec.DiskSizeGB > 0 {
		input.BlockDeviceMappings = []types.BlockDeviceMapping{
			{
				DeviceName: aws.String("/dev/sda1"),
				Ebs: &types.EbsBlockDevice{
					VolumeSize: aws.Int32(spec.DiskSizeGB),
					VolumeType: types.VolumeTypeGp3,
				},
			},
		}
	}
	if spec.UserData != "" {
		// AWS expects base64-encoded user-data
		encoded := base64.StdEncoding.EncodeToString([]byte(spec.UserData))
		input.UserData = aws.String(encoded)
	}

	// Add Name tag
	if spec.Name != "" {
		input.TagSpecifications = []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInstance,
				Tags: []types.Tag{
					{Key: aws.String("Name"), Value: aws.String(spec.Name)},
				},
			},
		}
	}

	out, err := client.RunInstances(context.Background(), input)
	if err != nil {
		return nil, err
	}

	if len(out.Instances) == 0 {
		return nil, fmt.Errorf("no instance returned")
	}

	i := out.Instances[0]
	inst := &Instance{
		ID:               aws.ToString(i.InstanceId),
		Name:             spec.Name,
		State:            string(i.State.Name),
		Type:             string(i.InstanceType),
		PrivateIP:        aws.ToString(i.PrivateIpAddress),
		PublicIP:         aws.ToString(i.PublicIpAddress),
		VPCID:            aws.ToString(i.VpcId),
		SubnetID:         aws.ToString(i.SubnetId),
		AvailabilityZone: aws.ToString(i.Placement.AvailabilityZone),
		KeyName:          aws.ToString(i.KeyName),
	}
	if i.LaunchTime != nil {
		inst.LaunchTime = i.LaunchTime.Format("2006-01-02 15:04")
	}
	return inst, nil
}

func (s *service) GetInstance(creds Credentials, instanceID string) (*Instance, error) {
	client := s.newEC2Client(creds)
	out, err := client.DescribeInstances(context.Background(), &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
	})
	if err != nil {
		return nil, err
	}
	for _, r := range out.Reservations {
		for _, i := range r.Instances {
			inst := Instance{
				ID:               aws.ToString(i.InstanceId),
				Name:             nameFromTags(i.Tags),
				State:            string(i.State.Name),
				Type:             string(i.InstanceType),
				PrivateIP:        aws.ToString(i.PrivateIpAddress),
				PublicIP:         aws.ToString(i.PublicIpAddress),
				VPCID:            aws.ToString(i.VpcId),
				SubnetID:         aws.ToString(i.SubnetId),
				AvailabilityZone: aws.ToString(i.Placement.AvailabilityZone),
				LaunchTime:       i.LaunchTime.Format("2006-01-02 15:04"),
				KeyName:          aws.ToString(i.KeyName),
			}
			return &inst, nil
		}
	}
	return nil, fmt.Errorf("instance %s not found", instanceID)
}

func (s *service) StartInstance(creds Credentials, instanceID string) error {
	client := s.newEC2Client(creds)
	_, err := client.StartInstances(context.Background(), &ec2.StartInstancesInput{
		InstanceIds: []string{instanceID},
	})
	return err
}

func (s *service) StopInstance(creds Credentials, instanceID string) error {
	client := s.newEC2Client(creds)
	_, err := client.StopInstances(context.Background(), &ec2.StopInstancesInput{
		InstanceIds: []string{instanceID},
	})
	return err
}

func (s *service) RebootInstance(creds Credentials, instanceID string) error {
	client := s.newEC2Client(creds)
	_, err := client.RebootInstances(context.Background(), &ec2.RebootInstancesInput{
		InstanceIds: []string{instanceID},
	})
	return err
}

func (s *service) TerminateInstance(creds Credentials, instanceID string) error {
	client := s.newEC2Client(creds)
	_, err := client.TerminateInstances(context.Background(), &ec2.TerminateInstancesInput{
		InstanceIds: []string{instanceID},
	})
	return err
}

func (s *service) CreateSecurityGroup(creds Credentials, spec SGSpec) (*SecurityGroup, error) {
	client := s.newEC2Client(creds)

	// Create the security group
	createOut, err := client.CreateSecurityGroup(context.Background(), &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(spec.Name),
		Description: aws.String(spec.Description),
		VpcId:       aws.String(spec.VPCID),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeSecurityGroup,
				Tags: []types.Tag{
					{Key: aws.String("Name"), Value: aws.String(spec.Name)},
					{Key: aws.String("managed-by"), Value: aws.String("palette-agentic")},
				},
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("create security group: %w", err)
	}

	sgID := aws.ToString(createOut.GroupId)

	// Add ingress rules
	var ingressRules []types.IpPermission
	for _, r := range spec.Rules {
		if r.Direction != "ingress" {
			continue
		}
		cidr := r.CIDR
		if cidr == "" {
			cidr = "0.0.0.0/0"
		}
		perm := types.IpPermission{
			IpProtocol: aws.String(r.Protocol),
			FromPort:   aws.Int32(r.FromPort),
			ToPort:     aws.Int32(r.ToPort),
			IpRanges:   []types.IpRange{{CidrIp: aws.String(cidr)}},
		}
		ingressRules = append(ingressRules, perm)
	}
	if len(ingressRules) > 0 {
		_, err = client.AuthorizeSecurityGroupIngress(context.Background(), &ec2.AuthorizeSecurityGroupIngressInput{
			GroupId:       aws.String(sgID),
			IpPermissions: ingressRules,
		})
		if err != nil {
			return nil, fmt.Errorf("add ingress rules: %w", err)
		}
	}

	// Add egress rules (beyond the default allow-all)
	var egressRules []types.IpPermission
	for _, r := range spec.Rules {
		if r.Direction != "egress" {
			continue
		}
		cidr := r.CIDR
		if cidr == "" {
			cidr = "0.0.0.0/0"
		}
		perm := types.IpPermission{
			IpProtocol: aws.String(r.Protocol),
			FromPort:   aws.Int32(r.FromPort),
			ToPort:     aws.Int32(r.ToPort),
			IpRanges:   []types.IpRange{{CidrIp: aws.String(cidr)}},
		}
		egressRules = append(egressRules, perm)
	}
	if len(egressRules) > 0 {
		_, err = client.AuthorizeSecurityGroupEgress(context.Background(), &ec2.AuthorizeSecurityGroupEgressInput{
			GroupId:       aws.String(sgID),
			IpPermissions: egressRules,
		})
		if err != nil {
			// Egress might fail if default allow-all already covers it — not fatal
			fmt.Printf("warning: add egress rules: %v\n", err)
		}
	}

	return &SecurityGroup{
		ID:          sgID,
		Name:        spec.Name,
		Description: spec.Description,
		VPCID:       spec.VPCID,
	}, nil
}

func (s *service) CreateNLB(creds Credentials, spec NLBSpec) (*NLBResult, error) {
	client := s.newELBClient(creds)
	ctx := context.Background()

	port := spec.Port
	if port == 0 {
		port = 6443
	}

	// 1. Create target group
	tgName := spec.Name + "-tg"
	tgOut, err := client.CreateTargetGroup(ctx, &elbv2.CreateTargetGroupInput{
		Name:       aws.String(tgName),
		Protocol:   elbtypes.ProtocolEnumTcp,
		Port:       aws.Int32(port),
		VpcId:      aws.String(spec.VPCID),
		TargetType: elbtypes.TargetTypeEnumInstance,
		HealthCheckProtocol: elbtypes.ProtocolEnumTcp,
		HealthCheckPort:     aws.String("traffic-port"),
		Tags: []elbtypes.Tag{
			{Key: aws.String("Name"), Value: aws.String(tgName)},
			{Key: aws.String("managed-by"), Value: aws.String("palette-agentic")},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("create target group: %w", err)
	}
	if len(tgOut.TargetGroups) == 0 {
		return nil, fmt.Errorf("no target group returned")
	}
	tgARN := aws.ToString(tgOut.TargetGroups[0].TargetGroupArn)

	// 2. Register targets (EC2 instances)
	if len(spec.TargetIDs) > 0 {
		targets := make([]elbtypes.TargetDescription, len(spec.TargetIDs))
		for i, id := range spec.TargetIDs {
			targets[i] = elbtypes.TargetDescription{
				Id:   aws.String(id),
				Port: aws.Int32(port),
			}
		}
		_, err = client.RegisterTargets(ctx, &elbv2.RegisterTargetsInput{
			TargetGroupArn: aws.String(tgARN),
			Targets:        targets,
		})
		if err != nil {
			return nil, fmt.Errorf("register targets: %w", err)
		}
	}

	// 3. Build subnet mappings
	var subnetMappings []elbtypes.SubnetMapping
	for _, subnetID := range spec.SubnetIDs {
		subnetMappings = append(subnetMappings, elbtypes.SubnetMapping{
			SubnetId: aws.String(subnetID),
		})
	}

	// 4. Create NLB
	lbOut, err := client.CreateLoadBalancer(ctx, &elbv2.CreateLoadBalancerInput{
		Name:           aws.String(spec.Name),
		Type:           elbtypes.LoadBalancerTypeEnumNetwork,
		Scheme:         elbtypes.LoadBalancerSchemeEnumInternetFacing,
		SubnetMappings: subnetMappings,
		Tags: []elbtypes.Tag{
			{Key: aws.String("Name"), Value: aws.String(spec.Name)},
			{Key: aws.String("managed-by"), Value: aws.String("palette-agentic")},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("create load balancer: %w", err)
	}
	if len(lbOut.LoadBalancers) == 0 {
		return nil, fmt.Errorf("no load balancer returned")
	}
	lb := lbOut.LoadBalancers[0]
	lbARN := aws.ToString(lb.LoadBalancerArn)
	lbDNS := aws.ToString(lb.DNSName)

	// 5. Create listener (forwards port to target group)
	_, err = client.CreateListener(ctx, &elbv2.CreateListenerInput{
		LoadBalancerArn: aws.String(lbARN),
		Protocol:        elbtypes.ProtocolEnumTcp,
		Port:            aws.Int32(port),
		DefaultActions: []elbtypes.Action{
			{
				Type:           elbtypes.ActionTypeEnumForward,
				TargetGroupArn: aws.String(tgARN),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("create listener: %w", err)
	}

	return &NLBResult{
		LBARN:       lbARN,
		LBName:      spec.Name,
		LBDNS:       lbDNS,
		TGARN:       tgARN,
		TGName:      tgName,
		Port:        port,
		TargetCount: len(spec.TargetIDs),
	}, nil
}

func (s *service) DeleteNLB(creds Credentials, arn string) error {
	client := s.newELBClient(creds)
	_, err := client.DeleteLoadBalancer(context.Background(), &elbv2.DeleteLoadBalancerInput{
		LoadBalancerArn: aws.String(arn),
	})
	return err
}

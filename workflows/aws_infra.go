package workflows

import (
	"encoding/json"

	"github.com/spectrocloud/palette-sdk-go/services/awsinfra"
)

// AWSInfra queries AWS infrastructure using credentials from env vars.
type AWSInfra struct {
	AWS awsinfra.Service
}

// AWSInfraInput is the input for AWS infra queries.
type AWSInfraInput struct {
	Mode       string               `json:"mode"`
	Region     string               `json:"region,omitempty"`
	VPCID      string               `json:"vpc_id,omitempty"`
	InstanceID string               `json:"instance_id,omitempty"`
	Filters    map[string]string    `json:"filters,omitempty"`
	Spec       *awsinfra.LaunchSpec `json:"spec,omitempty"`
	SGSpec     *awsinfra.SGSpec     `json:"sg_spec,omitempty"`
	NLBSpec    *awsinfra.NLBSpec    `json:"nlb_spec,omitempty"`
	ARN        string               `json:"arn,omitempty"`
}

// Run executes an AWS infra query.
func (w *AWSInfra) Run(input json.RawMessage) Result {
	var in AWSInfraInput
	if err := json.Unmarshal(input, &in); err != nil {
		return Errorf("invalid input: %v", err)
	}

	creds, err := awsinfra.ResolveCredentials(in.Region)
	if err != nil {
		return Errorf("AWS credentials: %v", err)
	}

	switch in.Mode {
	case "list_vpcs":
		vpcs, err := w.AWS.ListVPCs(creds)
		if err != nil {
			return Errorf("list_vpcs failed: %v", err)
		}
		return OK("list_vpcs completed.", map[string]any{"vpcs": vpcs})

	case "list_subnets":
		subnets, err := w.AWS.ListSubnets(creds, in.VPCID)
		if err != nil {
			return Errorf("list_subnets failed: %v", err)
		}
		return OK("list_subnets completed.", map[string]any{"subnets": subnets})

	case "list_key_pairs":
		pairs, err := w.AWS.ListKeyPairs(creds)
		if err != nil {
			return Errorf("list_key_pairs failed: %v", err)
		}
		return OK("list_key_pairs completed.", map[string]any{"key_pairs": pairs})

	case "list_security_groups":
		groups, err := w.AWS.ListSecurityGroups(creds, in.VPCID)
		if err != nil {
			return Errorf("list_security_groups failed: %v", err)
		}
		return OK("list_security_groups completed.", map[string]any{"security_groups": groups})

	case "list_availability_zones":
		zones, err := w.AWS.ListAvailabilityZones(creds)
		if err != nil {
			return Errorf("list_availability_zones failed: %v", err)
		}
		return OK("list_availability_zones completed.", map[string]any{"zones": zones})

	case "list_instance_types":
		types, err := w.AWS.ListInstanceTypes(creds)
		if err != nil {
			return Errorf("list_instance_types failed: %v", err)
		}
		return OK("list_instance_types completed.", map[string]any{"instance_types": types})

	case "list_eks_clusters":
		clusters, err := w.AWS.ListEKSClusters(creds)
		if err != nil {
			return Errorf("list_eks_clusters failed: %v", err)
		}
		return OK("list_eks_clusters completed.", map[string]any{"clusters": clusters})

	case "create_nlb":
		if in.NLBSpec == nil {
			return Errorf("nlb_spec is required for create_nlb")
		}
		nlb, err := w.AWS.CreateNLB(creds, *in.NLBSpec)
		if err != nil {
			return Errorf("create_nlb failed: %v", err)
		}
		return OK("NLB created.", map[string]any{"nlb": nlb})

	case "delete_nlb":
		if in.ARN == "" {
			return Errorf("arn is required for delete_nlb")
		}
		if err := w.AWS.DeleteNLB(creds, in.ARN); err != nil {
			return Errorf("delete_nlb failed: %v", err)
		}
		return OK("NLB deleted.", map[string]any{"arn": in.ARN})

	case "create_security_group":
		if in.SGSpec == nil {
			return Errorf("sg_spec is required for create_security_group")
		}
		sg, err := w.AWS.CreateSecurityGroup(creds, *in.SGSpec)
		if err != nil {
			return Errorf("create_security_group failed: %v", err)
		}
		return OK("Security group created.", map[string]any{"security_group": sg})

	case "launch_instance":
		if in.Spec == nil {
			return Errorf("spec is required for launch_instance")
		}
		if in.Spec.ImageID == "" {
			return Errorf("spec.imageId (AMI ID) is required")
		}
		if in.Spec.InstanceType == "" {
			in.Spec.InstanceType = "t3.medium"
		}
		inst, err := w.AWS.LaunchInstance(creds, *in.Spec)
		if err != nil {
			return Errorf("launch_instance failed: %v", err)
		}
		return OK("Instance launched.", map[string]any{"instance": inst})

	case "list_instances":
		instances, err := w.AWS.ListInstances(creds, in.Filters)
		if err != nil {
			return Errorf("list_instances failed: %v", err)
		}
		return OK("list_instances completed.", map[string]any{"instances": instances})

	case "get_instance":
		if in.InstanceID == "" {
			return Errorf("instance_id is required")
		}
		inst, err := w.AWS.GetInstance(creds, in.InstanceID)
		if err != nil {
			return Errorf("get_instance failed: %v", err)
		}
		return OK("get_instance completed.", map[string]any{"instance": inst})

	case "start_instance":
		if in.InstanceID == "" {
			return Errorf("instance_id is required")
		}
		if err := w.AWS.StartInstance(creds, in.InstanceID); err != nil {
			return Errorf("start_instance failed: %v", err)
		}
		return OK("Instance started.", map[string]any{"instance_id": in.InstanceID})

	case "stop_instance":
		if in.InstanceID == "" {
			return Errorf("instance_id is required")
		}
		if err := w.AWS.StopInstance(creds, in.InstanceID); err != nil {
			return Errorf("stop_instance failed: %v", err)
		}
		return OK("Instance stopped.", map[string]any{"instance_id": in.InstanceID})

	case "reboot_instance":
		if in.InstanceID == "" {
			return Errorf("instance_id is required")
		}
		if err := w.AWS.RebootInstance(creds, in.InstanceID); err != nil {
			return Errorf("reboot_instance failed: %v", err)
		}
		return OK("Instance rebooted.", map[string]any{"instance_id": in.InstanceID})

	case "terminate_instance":
		if in.InstanceID == "" {
			return Errorf("instance_id is required")
		}
		if err := w.AWS.TerminateInstance(creds, in.InstanceID); err != nil {
			return Errorf("terminate_instance failed: %v", err)
		}
		return OK("Instance terminated.", map[string]any{"instance_id": in.InstanceID})

	default:
		return Errorf("unknown mode %q", in.Mode)
	}
}

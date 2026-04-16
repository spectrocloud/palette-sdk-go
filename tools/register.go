package tools

import (
	"encoding/json"

	"github.com/spectrocloud/palette-sdk-go/services/access"
	"github.com/spectrocloud/palette-sdk-go/services/awsinfra"
	"github.com/spectrocloud/palette-sdk-go/services/cloudaccounts"
	"github.com/spectrocloud/palette-sdk-go/services/clusters"
	"github.com/spectrocloud/palette-sdk-go/services/diagnostics"
	"github.com/spectrocloud/palette-sdk-go/services/edge"
	"github.com/spectrocloud/palette-sdk-go/services/platform"
	"github.com/spectrocloud/palette-sdk-go/services/profiles"
	"github.com/spectrocloud/palette-sdk-go/services/registries"
	"github.com/spectrocloud/palette-sdk-go/services/vmo"
	"github.com/spectrocloud/palette-sdk-go/workflows"
	"github.com/spectrocloud/palette-sdk-go/workflows/generated"
)

// Services holds all service dependencies needed to register workflow tools.
type Services struct {
	Clusters      clusters.Service
	CloudAccounts cloudaccounts.Service
	Profiles      profiles.Service
	Access        access.Service
	Registries    registries.Service
	Diagnostics   diagnostics.Service
	Platform      platform.Service
	VMO           vmo.Service
	AWSInfra      awsinfra.Service
	Edge          edge.Service
}

// RegisterAll registers all generated and hand-written workflow tools.
func RegisterAll(r *Registry, svc Services) {
	// Generated tools
	genSvc := generated.Services{
		Clusters:      svc.Clusters,
		CloudAccounts: svc.CloudAccounts,
		Profiles:      svc.Profiles,
		Access:        svc.Access,
		Registries:    svc.Registries,
		Diagnostics:   svc.Diagnostics,
		Platform:      svc.Platform,
		VMO:           svc.VMO,
	}
	for _, t := range generated.GetAll(genSvc) {
		r.Register(Tool{
			Name:        t.Name,
			Description: t.Description,
			Run:         t.Run,
		})
	}

	// Hand-written: AWS Infrastructure
	if svc.AWSInfra != nil {
		wf := &workflows.AWSInfra{AWS: svc.AWSInfra}
		r.Register(Tool{
			Name:        "aws_infra_api",
			Description: "Query and manage AWS infrastructure — VPCs, subnets, EC2 instances, key pairs",
			Run:         func(input json.RawMessage) string { return wf.Run(input).String() },
		})
	}

	// Hand-written: Edge hosts and registration tokens
	if svc.Edge != nil {
		wf := &workflows.Edge{Edge: svc.Edge}
		r.Register(Tool{
			Name:        "edge_api",
			Description: "Manage edge hosts and registration tokens — list hosts, create tokens, query by tags",
			Run:         func(input json.RawMessage) string { return wf.Run(input).String() },
		})
	}
}

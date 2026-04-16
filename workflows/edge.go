package workflows

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/services/edge"
)

// Edge handles edge host and registration token operations.
type Edge struct {
	Edge edge.Service
}

// EdgeInput is the input for edge operations.
type EdgeInput struct {
	Mode      string            `json:"mode"`
	Name      string            `json:"name,omitempty"`
	UID       string            `json:"uid,omitempty"`
	Tags      map[string]string `json:"tags,omitempty"`
	TokenName string            `json:"token_name,omitempty"`
	HostNames []string          `json:"host_names,omitempty"` // For batch create
}

// Run executes an edge operation.
func (w *Edge) Run(input json.RawMessage) Result {
	var in EdgeInput
	if err := json.Unmarshal(input, &in); err != nil {
		return Errorf("invalid input: %v", err)
	}

	switch in.Mode {
	// Token operations
	case "get_token":
		return w.getToken(in)
	case "create_token":
		return w.createToken(in)
	case "delete_token":
		return w.deleteToken(in)

	// Edge host operations
	case "create_edge_host":
		return w.createEdgeHost(in)
	case "list_edge_hosts":
		return w.listEdgeHosts()
	case "get_edge_host":
		return w.getEdgeHost(in)
	case "get_edge_hosts_by_tags":
		return w.getEdgeHostsByTags(in)
	case "delete_edge_host":
		return w.deleteEdgeHost(in)

	default:
		return Errorf("unknown mode %q", in.Mode)
	}
}

func (w *Edge) getToken(in EdgeInput) Result {
	name := in.TokenName
	if name == "" {
		name = in.Name
	}
	if name == "" {
		return Errorf("token_name or name is required")
	}
	token, err := w.Edge.GetTokenByName(name)
	if err != nil {
		return Errorf("get_token failed: %v", err)
	}
	return OK("Token retrieved.", map[string]any{"token": token})
}

func (w *Edge) createToken(in EdgeInput) Result {
	name := in.TokenName
	if name == "" {
		name = in.Name
	}
	if name == "" {
		name = "edge-registration-token"
	}

	// Check for an existing active, non-expired token with the same name.
	existing, err := w.Edge.GetTokenByName(name)
	if err == nil && existing != nil &&
		existing.Status != nil && existing.Status.IsActive &&
		existing.Spec != nil {
		expiry := time.Time(existing.Spec.Expiry)
		if !expiry.IsZero() && expiry.After(time.Now()) {
			// Reuse the existing token.
			tokenValue, err := w.Edge.GetTokenValue(name)
			if err == nil && tokenValue != "" {
				uid := ""
				if existing.Metadata != nil {
					uid = existing.Metadata.UID
				}
				return OK(fmt.Sprintf("Reusing existing token '%s' (expires %s).", name, expiry.Format("2006-01-02")), map[string]any{
					"token_uid":   uid,
					"token_value": tokenValue,
					"token_name":  name,
					"reused":      true,
				})
			}
		}
	}

	// No valid token found — create a new one.
	expiry := models.V1Time(strfmt.DateTime(time.Now().AddDate(1, 0, 0)))
	body := &models.V1EdgeTokenEntity{
		Metadata: &models.V1ObjectMeta{
			Name: name,
		},
		Spec: &models.V1EdgeTokenSpecEntity{
			Expiry: expiry,
		},
	}
	uid, tokenValue, err := w.Edge.CreateToken(name, body)
	if err != nil {
		return Errorf("create_token failed: %v", err)
	}
	return OK(fmt.Sprintf("Token '%s' created.", name), map[string]any{
		"token_uid":   uid,
		"token_value": tokenValue,
		"token_name":  name,
	})
}

func (w *Edge) deleteToken(in EdgeInput) Result {
	if in.UID == "" {
		return Errorf("uid is required for delete_token")
	}
	if err := w.Edge.DeleteToken(in.UID); err != nil {
		return Errorf("delete_token failed: %v", err)
	}
	return OK("Token deleted.", map[string]any{"token_uid": in.UID})
}

func (w *Edge) createEdgeHost(in EdgeInput) Result {
	// Support batch creation via host_names or single via name
	names := in.HostNames
	if len(names) == 0 && in.Name != "" {
		names = []string{in.Name}
	}
	if len(names) == 0 {
		return Errorf("name or host_names is required for create_edge_host")
	}

	created := make([]map[string]string, 0, len(names))
	for _, name := range names {
		body := &models.V1EdgeHostDeviceEntity{
			Metadata: &models.V1ObjectTagsEntity{
				Name:   name,
				UID:    name, // UID matches name so VMs pair by site name
				Labels: in.Tags,
			},
		}
		uid, err := w.Edge.CreateEdgeHost(body)
		if err != nil {
			return Errorf("create_edge_host '%s' failed: %v", name, err)
		}
		created = append(created, map[string]string{"name": name, "uid": uid})
	}

	return OK(fmt.Sprintf("Created %d edge host(s).", len(created)), map[string]any{
		"edge_hosts": created,
	})
}

func (w *Edge) listEdgeHosts() Result {
	hosts, err := w.Edge.ListEdgeHosts()
	if err != nil {
		return Errorf("list_edge_hosts failed: %v", err)
	}
	return OK(fmt.Sprintf("Found %d edge hosts.", len(hosts)), map[string]any{"edge_hosts": hosts})
}

func (w *Edge) getEdgeHost(in EdgeInput) Result {
	if in.UID != "" {
		host, err := w.Edge.GetEdgeHost(in.UID)
		if err != nil {
			return Errorf("get_edge_host failed: %v", err)
		}
		return OK("Edge host retrieved.", map[string]any{"edge_host": host})
	}
	if in.Name != "" {
		host, err := w.Edge.GetEdgeHostByName(in.Name)
		if err != nil {
			return Errorf("get_edge_host failed: %v", err)
		}
		return OK("Edge host retrieved.", map[string]any{"edge_host": host})
	}
	return Errorf("uid or name is required")
}

func (w *Edge) getEdgeHostsByTags(in EdgeInput) Result {
	if len(in.Tags) == 0 {
		return Errorf("tags are required for get_edge_hosts_by_tags")
	}
	hosts, err := w.Edge.GetEdgeHostsByTags(in.Tags)
	if err != nil {
		return Errorf("get_edge_hosts_by_tags failed: %v", err)
	}
	return OK(fmt.Sprintf("Found %d edge hosts matching tags.", len(hosts)), map[string]any{"edge_hosts": hosts})
}

func (w *Edge) deleteEdgeHost(in EdgeInput) Result {
	if in.UID == "" {
		return Errorf("uid is required for delete_edge_host")
	}
	if err := w.Edge.DeleteEdgeHost(in.UID); err != nil {
		return Errorf("delete_edge_host failed: %v", err)
	}
	return OK("Edge host deleted.", map[string]any{"edge_host_uid": in.UID})
}

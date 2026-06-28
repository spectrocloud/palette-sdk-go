package client

import (
	"context"
	"time"

	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
	"github.com/spectrocloud/palette-sdk-go/client/apiutil"
)

type profileMetaPatchBody struct {
	body *models.V1ProfileMetaEntity
}

func (b profileMetaPatchBody) MarshalJSON() ([]byte, error) {
	return apiutil.MarshalProfileMetaPatchBody(b.body)
}

type clusterProfileMetadataPatchParams struct {
	uid     string
	body    *models.V1ProfileMetaEntity
	timeout time.Duration
	Context context.Context
}

func newClusterProfileMetadataPatchParams(ctx context.Context, uid string, body *models.V1ProfileMetaEntity) *clusterProfileMetadataPatchParams {
	return &clusterProfileMetadataPatchParams{
		uid:     uid,
		body:    body,
		timeout: cr.DefaultTimeout,
		Context: ctx,
	}
}

func (o *clusterProfileMetadataPatchParams) WriteToRequest(r runtime.ClientRequest, _ strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	if o.body != nil {
		if err := r.SetBodyParam(profileMetaPatchBody{body: o.body}); err != nil {
			return err
		}
	}
	return r.SetPathParam("uid", o.uid)
}

func (h *V1Client) patchClusterProfileMetadata(uid string, metadata *models.V1ProfileMetaEntity) error {
	params := newClusterProfileMetadataPatchParams(h.ctx, uid, metadata)
	_, err := h.getTransport().Submit(&runtime.ClientOperation{
		ID:                 "v1ClusterProfilesUidMetadataUpdate",
		Method:             "PATCH",
		PathPattern:        "/v1/clusterprofiles/{uid}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            h.schemes,
		Params:             params,
		Reader:             clusterProfileMetadataPatchReader{},
		Context:            params.Context,
	})
	return err
}

type clusterProfileMetadataPatchReader struct{}

func (clusterProfileMetadataPatchReader) ReadResponse(response runtime.ClientResponse, _ runtime.Consumer) (interface{}, error) {
	if response.Code() == 204 {
		return struct{}{}, nil
	}
	return nil, runtime.NewAPIError(
		"response status code does not match any response statuses defined for this endpoint in the swagger spec",
		response,
		response.Code(),
	)
}

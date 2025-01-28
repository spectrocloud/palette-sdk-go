// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1ClusterProfilesUIDMetadataUpdateReader is a Reader for the V1ClusterProfilesUIDMetadataUpdate structure.
type V1ClusterProfilesUIDMetadataUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterProfilesUIDMetadataUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1ClusterProfilesUIDMetadataUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterProfilesUIDMetadataUpdateNoContent creates a V1ClusterProfilesUIDMetadataUpdateNoContent with default headers values
func NewV1ClusterProfilesUIDMetadataUpdateNoContent() *V1ClusterProfilesUIDMetadataUpdateNoContent {
	return &V1ClusterProfilesUIDMetadataUpdateNoContent{}
}

/*V1ClusterProfilesUIDMetadataUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1ClusterProfilesUIDMetadataUpdateNoContent struct {
}

func (o *V1ClusterProfilesUIDMetadataUpdateNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusterprofiles/{uid}/metadata][%d] v1ClusterProfilesUidMetadataUpdateNoContent ", 204)
}

func (o *V1ClusterProfilesUIDMetadataUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

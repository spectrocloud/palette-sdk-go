// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1ClusterProfilesMetadataReader is a Reader for the V1ClusterProfilesMetadata structure.
type V1ClusterProfilesMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterProfilesMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterProfilesMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterProfilesMetadataOK creates a V1ClusterProfilesMetadataOK with default headers values
func NewV1ClusterProfilesMetadataOK() *V1ClusterProfilesMetadataOK {
	return &V1ClusterProfilesMetadataOK{}
}

/*V1ClusterProfilesMetadataOK handles this case with default header values.

An array of cluster summary items
*/
type V1ClusterProfilesMetadataOK struct {
	Payload *models.V1ClusterProfilesMetadata
}

func (o *V1ClusterProfilesMetadataOK) Error() string {
	return fmt.Sprintf("[GET /v1/dashboard/clusterprofiles/metadata][%d] v1ClusterProfilesMetadataOK  %+v", 200, o.Payload)
}

func (o *V1ClusterProfilesMetadataOK) GetPayload() *models.V1ClusterProfilesMetadata {
	return o.Payload
}

func (o *V1ClusterProfilesMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterProfilesMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

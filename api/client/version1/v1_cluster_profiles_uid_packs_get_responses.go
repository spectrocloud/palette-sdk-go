// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1ClusterProfilesUIDPacksGetReader is a Reader for the V1ClusterProfilesUIDPacksGet structure.
type V1ClusterProfilesUIDPacksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterProfilesUIDPacksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterProfilesUIDPacksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterProfilesUIDPacksGetOK creates a V1ClusterProfilesUIDPacksGetOK with default headers values
func NewV1ClusterProfilesUIDPacksGetOK() *V1ClusterProfilesUIDPacksGetOK {
	return &V1ClusterProfilesUIDPacksGetOK{}
}

/*
V1ClusterProfilesUIDPacksGetOK handles this case with default header values.

OK
*/
type V1ClusterProfilesUIDPacksGetOK struct {
	Payload *models.V1ClusterProfilePacksEntities
}

func (o *V1ClusterProfilesUIDPacksGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusterprofiles/{uid}/packs][%d] v1ClusterProfilesUidPacksGetOK  %+v", 200, o.Payload)
}

func (o *V1ClusterProfilesUIDPacksGetOK) GetPayload() *models.V1ClusterProfilePacksEntities {
	return o.Payload
}

func (o *V1ClusterProfilesUIDPacksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterProfilePacksEntities)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

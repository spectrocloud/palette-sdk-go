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

// V1OverlordsUIDMaasClusterProfileReader is a Reader for the V1OverlordsUIDMaasClusterProfile structure.
type V1OverlordsUIDMaasClusterProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OverlordsUIDMaasClusterProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OverlordsUIDMaasClusterProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OverlordsUIDMaasClusterProfileOK creates a V1OverlordsUIDMaasClusterProfileOK with default headers values
func NewV1OverlordsUIDMaasClusterProfileOK() *V1OverlordsUIDMaasClusterProfileOK {
	return &V1OverlordsUIDMaasClusterProfileOK{}
}

/*
V1OverlordsUIDMaasClusterProfileOK handles this case with default header values.

OK
*/
type V1OverlordsUIDMaasClusterProfileOK struct {
	Payload *models.V1ClusterProfile
}

func (o *V1OverlordsUIDMaasClusterProfileOK) Error() string {
	return fmt.Sprintf("[GET /v1/overlords/maas/{uid}/clusterprofile][%d] v1OverlordsUidMaasClusterProfileOK  %+v", 200, o.Payload)
}

func (o *V1OverlordsUIDMaasClusterProfileOK) GetPayload() *models.V1ClusterProfile {
	return o.Payload
}

func (o *V1OverlordsUIDMaasClusterProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

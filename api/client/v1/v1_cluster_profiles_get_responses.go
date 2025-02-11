// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1ClusterProfilesGetReader is a Reader for the V1ClusterProfilesGet structure.
type V1ClusterProfilesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterProfilesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterProfilesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterProfilesGetOK creates a V1ClusterProfilesGetOK with default headers values
func NewV1ClusterProfilesGetOK() *V1ClusterProfilesGetOK {
	return &V1ClusterProfilesGetOK{}
}

/*
V1ClusterProfilesGetOK handles this case with default header values.

OK
*/
type V1ClusterProfilesGetOK struct {
	Payload *models.V1ClusterProfile
}

func (o *V1ClusterProfilesGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusterprofiles/{uid}][%d] v1ClusterProfilesGetOK  %+v", 200, o.Payload)
}

func (o *V1ClusterProfilesGetOK) GetPayload() *models.V1ClusterProfile {
	return o.Payload
}

func (o *V1ClusterProfilesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1OpenStackProjectsGetReader is a Reader for the V1OpenStackProjectsGet structure.
type V1OpenStackProjectsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OpenStackProjectsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OpenStackProjectsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OpenStackProjectsGetOK creates a V1OpenStackProjectsGetOK with default headers values
func NewV1OpenStackProjectsGetOK() *V1OpenStackProjectsGetOK {
	return &V1OpenStackProjectsGetOK{}
}

/*
V1OpenStackProjectsGetOK handles this case with default header values.

(empty)
*/
type V1OpenStackProjectsGetOK struct {
	Payload *models.V1OpenStackProjects
}

func (o *V1OpenStackProjectsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/openstack/projects][%d] v1OpenStackProjectsGetOK  %+v", 200, o.Payload)
}

func (o *V1OpenStackProjectsGetOK) GetPayload() *models.V1OpenStackProjects {
	return o.Payload
}

func (o *V1OpenStackProjectsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OpenStackProjects)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1OpenstackAccountsUIDProjectsReader is a Reader for the V1OpenstackAccountsUIDProjects structure.
type V1OpenstackAccountsUIDProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OpenstackAccountsUIDProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OpenstackAccountsUIDProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OpenstackAccountsUIDProjectsOK creates a V1OpenstackAccountsUIDProjectsOK with default headers values
func NewV1OpenstackAccountsUIDProjectsOK() *V1OpenstackAccountsUIDProjectsOK {
	return &V1OpenstackAccountsUIDProjectsOK{}
}

/*
V1OpenstackAccountsUIDProjectsOK handles this case with default header values.

(empty)
*/
type V1OpenstackAccountsUIDProjectsOK struct {
	Payload *models.V1OpenStackProjects
}

func (o *V1OpenstackAccountsUIDProjectsOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/openstack/{uid}/properties/projects][%d] v1OpenstackAccountsUidProjectsOK  %+v", 200, o.Payload)
}

func (o *V1OpenstackAccountsUIDProjectsOK) GetPayload() *models.V1OpenStackProjects {
	return o.Payload
}

func (o *V1OpenstackAccountsUIDProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OpenStackProjects)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

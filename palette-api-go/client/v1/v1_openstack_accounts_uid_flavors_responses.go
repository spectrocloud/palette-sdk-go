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

// V1OpenstackAccountsUIDFlavorsReader is a Reader for the V1OpenstackAccountsUIDFlavors structure.
type V1OpenstackAccountsUIDFlavorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OpenstackAccountsUIDFlavorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OpenstackAccountsUIDFlavorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OpenstackAccountsUIDFlavorsOK creates a V1OpenstackAccountsUIDFlavorsOK with default headers values
func NewV1OpenstackAccountsUIDFlavorsOK() *V1OpenstackAccountsUIDFlavorsOK {
	return &V1OpenstackAccountsUIDFlavorsOK{}
}

/*V1OpenstackAccountsUIDFlavorsOK handles this case with default header values.

(empty)
*/
type V1OpenstackAccountsUIDFlavorsOK struct {
	Payload *models.V1OpenStackFlavors
}

func (o *V1OpenstackAccountsUIDFlavorsOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/openstack/{uid}/properties/flavors][%d] v1OpenstackAccountsUidFlavorsOK  %+v", 200, o.Payload)
}

func (o *V1OpenstackAccountsUIDFlavorsOK) GetPayload() *models.V1OpenStackFlavors {
	return o.Payload
}

func (o *V1OpenstackAccountsUIDFlavorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OpenStackFlavors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

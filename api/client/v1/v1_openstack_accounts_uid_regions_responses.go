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

// V1OpenstackAccountsUIDRegionsReader is a Reader for the V1OpenstackAccountsUIDRegions structure.
type V1OpenstackAccountsUIDRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OpenstackAccountsUIDRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OpenstackAccountsUIDRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OpenstackAccountsUIDRegionsOK creates a V1OpenstackAccountsUIDRegionsOK with default headers values
func NewV1OpenstackAccountsUIDRegionsOK() *V1OpenstackAccountsUIDRegionsOK {
	return &V1OpenstackAccountsUIDRegionsOK{}
}

/*
V1OpenstackAccountsUIDRegionsOK handles this case with default header values.

(empty)
*/
type V1OpenstackAccountsUIDRegionsOK struct {
	Payload *models.V1OpenStackRegions
}

func (o *V1OpenstackAccountsUIDRegionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/openstack/{uid}/properties/regions][%d] v1OpenstackAccountsUidRegionsOK  %+v", 200, o.Payload)
}

func (o *V1OpenstackAccountsUIDRegionsOK) GetPayload() *models.V1OpenStackRegions {
	return o.Payload
}

func (o *V1OpenstackAccountsUIDRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OpenStackRegions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

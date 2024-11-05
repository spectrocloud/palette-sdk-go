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

// V1CloudAccountsVsphereListReader is a Reader for the V1CloudAccountsVsphereList structure.
type V1CloudAccountsVsphereListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsVsphereListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsVsphereListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsVsphereListOK creates a V1CloudAccountsVsphereListOK with default headers values
func NewV1CloudAccountsVsphereListOK() *V1CloudAccountsVsphereListOK {
	return &V1CloudAccountsVsphereListOK{}
}

/*
V1CloudAccountsVsphereListOK handles this case with default header values.

An array of cloud account items
*/
type V1CloudAccountsVsphereListOK struct {
	Payload *models.V1VsphereAccounts
}

func (o *V1CloudAccountsVsphereListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/vsphere][%d] v1CloudAccountsVsphereListOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsVsphereListOK) GetPayload() *models.V1VsphereAccounts {
	return o.Payload
}

func (o *V1CloudAccountsVsphereListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1VsphereAccounts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

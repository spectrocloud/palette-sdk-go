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

// V1CloudAccountsGcpListReader is a Reader for the V1CloudAccountsGcpList structure.
type V1CloudAccountsGcpListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsGcpListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsGcpListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsGcpListOK creates a V1CloudAccountsGcpListOK with default headers values
func NewV1CloudAccountsGcpListOK() *V1CloudAccountsGcpListOK {
	return &V1CloudAccountsGcpListOK{}
}

/*
V1CloudAccountsGcpListOK handles this case with default header values.

An array of gcp cloud account items
*/
type V1CloudAccountsGcpListOK struct {
	Payload *models.V1GcpAccounts
}

func (o *V1CloudAccountsGcpListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/gcp][%d] v1CloudAccountsGcpListOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsGcpListOK) GetPayload() *models.V1GcpAccounts {
	return o.Payload
}

func (o *V1CloudAccountsGcpListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GcpAccounts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1CloudAccountsOpenStackGetReader is a Reader for the V1CloudAccountsOpenStackGet structure.
type V1CloudAccountsOpenStackGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsOpenStackGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsOpenStackGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsOpenStackGetOK creates a V1CloudAccountsOpenStackGetOK with default headers values
func NewV1CloudAccountsOpenStackGetOK() *V1CloudAccountsOpenStackGetOK {
	return &V1CloudAccountsOpenStackGetOK{}
}

/*
V1CloudAccountsOpenStackGetOK handles this case with default header values.

OK
*/
type V1CloudAccountsOpenStackGetOK struct {
	Payload *models.V1OpenStackAccount
}

func (o *V1CloudAccountsOpenStackGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/openstack/{uid}][%d] v1CloudAccountsOpenStackGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsOpenStackGetOK) GetPayload() *models.V1OpenStackAccount {
	return o.Payload
}

func (o *V1CloudAccountsOpenStackGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OpenStackAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1CloudAccountsMaasListReader is a Reader for the V1CloudAccountsMaasList structure.
type V1CloudAccountsMaasListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsMaasListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsMaasListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsMaasListOK creates a V1CloudAccountsMaasListOK with default headers values
func NewV1CloudAccountsMaasListOK() *V1CloudAccountsMaasListOK {
	return &V1CloudAccountsMaasListOK{}
}

/*
V1CloudAccountsMaasListOK handles this case with default header values.

An array of cloud account items
*/
type V1CloudAccountsMaasListOK struct {
	Payload *models.V1MaasAccounts
}

func (o *V1CloudAccountsMaasListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/maas][%d] v1CloudAccountsMaasListOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsMaasListOK) GetPayload() *models.V1MaasAccounts {
	return o.Payload
}

func (o *V1CloudAccountsMaasListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MaasAccounts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

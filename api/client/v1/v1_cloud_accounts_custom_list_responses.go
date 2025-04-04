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

// V1CloudAccountsCustomListReader is a Reader for the V1CloudAccountsCustomList structure.
type V1CloudAccountsCustomListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsCustomListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsCustomListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsCustomListOK creates a V1CloudAccountsCustomListOK with default headers values
func NewV1CloudAccountsCustomListOK() *V1CloudAccountsCustomListOK {
	return &V1CloudAccountsCustomListOK{}
}

/*
V1CloudAccountsCustomListOK handles this case with default header values.

An array of cloud account by specified cloud type items
*/
type V1CloudAccountsCustomListOK struct {
	Payload *models.V1CustomAccounts
}

func (o *V1CloudAccountsCustomListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/cloudTypes/{cloudType}][%d] v1CloudAccountsCustomListOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsCustomListOK) GetPayload() *models.V1CustomAccounts {
	return o.Payload
}

func (o *V1CloudAccountsCustomListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CustomAccounts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1CloudAccountsCoxEdgeListReader is a Reader for the V1CloudAccountsCoxEdgeList structure.
type V1CloudAccountsCoxEdgeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsCoxEdgeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsCoxEdgeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsCoxEdgeListOK creates a V1CloudAccountsCoxEdgeListOK with default headers values
func NewV1CloudAccountsCoxEdgeListOK() *V1CloudAccountsCoxEdgeListOK {
	return &V1CloudAccountsCoxEdgeListOK{}
}

/*
V1CloudAccountsCoxEdgeListOK handles this case with default header values.

An array of cloud account items
*/
type V1CloudAccountsCoxEdgeListOK struct {
	Payload *models.V1CoxEdgeAccounts
}

func (o *V1CloudAccountsCoxEdgeListOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/coxedge][%d] v1CloudAccountsCoxEdgeListOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsCoxEdgeListOK) GetPayload() *models.V1CoxEdgeAccounts {
	return o.Payload
}

func (o *V1CloudAccountsCoxEdgeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CoxEdgeAccounts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

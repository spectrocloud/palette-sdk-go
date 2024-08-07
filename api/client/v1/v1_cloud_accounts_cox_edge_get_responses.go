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

// V1CloudAccountsCoxEdgeGetReader is a Reader for the V1CloudAccountsCoxEdgeGet structure.
type V1CloudAccountsCoxEdgeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudAccountsCoxEdgeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CloudAccountsCoxEdgeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudAccountsCoxEdgeGetOK creates a V1CloudAccountsCoxEdgeGetOK with default headers values
func NewV1CloudAccountsCoxEdgeGetOK() *V1CloudAccountsCoxEdgeGetOK {
	return &V1CloudAccountsCoxEdgeGetOK{}
}

/*
V1CloudAccountsCoxEdgeGetOK handles this case with default header values.

OK
*/
type V1CloudAccountsCoxEdgeGetOK struct {
	Payload *models.V1CoxEdgeAccount
}

func (o *V1CloudAccountsCoxEdgeGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/coxedge/{uid}][%d] v1CloudAccountsCoxEdgeGetOK  %+v", 200, o.Payload)
}

func (o *V1CloudAccountsCoxEdgeGetOK) GetPayload() *models.V1CoxEdgeAccount {
	return o.Payload
}

func (o *V1CloudAccountsCoxEdgeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CoxEdgeAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

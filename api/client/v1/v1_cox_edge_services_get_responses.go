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

// V1CoxEdgeServicesGetReader is a Reader for the V1CoxEdgeServicesGet structure.
type V1CoxEdgeServicesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CoxEdgeServicesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CoxEdgeServicesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CoxEdgeServicesGetOK creates a V1CoxEdgeServicesGetOK with default headers values
func NewV1CoxEdgeServicesGetOK() *V1CoxEdgeServicesGetOK {
	return &V1CoxEdgeServicesGetOK{}
}

/*
V1CoxEdgeServicesGetOK handles this case with default header values.

List of CoxEdge services
*/
type V1CoxEdgeServicesGetOK struct {
	Payload *models.V1CoxEdgeServices
}

func (o *V1CoxEdgeServicesGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/coxedge/services][%d] v1CoxEdgeServicesGetOK  %+v", 200, o.Payload)
}

func (o *V1CoxEdgeServicesGetOK) GetPayload() *models.V1CoxEdgeServices {
	return o.Payload
}

func (o *V1CoxEdgeServicesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CoxEdgeServices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

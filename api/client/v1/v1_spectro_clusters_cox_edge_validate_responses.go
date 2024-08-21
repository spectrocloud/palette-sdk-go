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

// V1SpectroClustersCoxEdgeValidateReader is a Reader for the V1SpectroClustersCoxEdgeValidate structure.
type V1SpectroClustersCoxEdgeValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersCoxEdgeValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersCoxEdgeValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersCoxEdgeValidateOK creates a V1SpectroClustersCoxEdgeValidateOK with default headers values
func NewV1SpectroClustersCoxEdgeValidateOK() *V1SpectroClustersCoxEdgeValidateOK {
	return &V1SpectroClustersCoxEdgeValidateOK{}
}

/*
V1SpectroClustersCoxEdgeValidateOK handles this case with default header values.

Azure Cluster validation response
*/
type V1SpectroClustersCoxEdgeValidateOK struct {
	Payload *models.V1SpectroClusterValidatorResponse
}

func (o *V1SpectroClustersCoxEdgeValidateOK) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/coxedge/validate][%d] v1SpectroClustersCoxEdgeValidateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersCoxEdgeValidateOK) GetPayload() *models.V1SpectroClusterValidatorResponse {
	return o.Payload
}

func (o *V1SpectroClustersCoxEdgeValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterValidatorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
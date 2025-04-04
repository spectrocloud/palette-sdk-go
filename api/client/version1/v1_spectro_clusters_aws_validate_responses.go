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

// V1SpectroClustersAwsValidateReader is a Reader for the V1SpectroClustersAwsValidate structure.
type V1SpectroClustersAwsValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersAwsValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersAwsValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersAwsValidateOK creates a V1SpectroClustersAwsValidateOK with default headers values
func NewV1SpectroClustersAwsValidateOK() *V1SpectroClustersAwsValidateOK {
	return &V1SpectroClustersAwsValidateOK{}
}

/*
V1SpectroClustersAwsValidateOK handles this case with default header values.

Aws Cluster validation response
*/
type V1SpectroClustersAwsValidateOK struct {
	Payload *models.V1SpectroClusterValidatorResponse
}

func (o *V1SpectroClustersAwsValidateOK) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/aws/validate][%d] v1SpectroClustersAwsValidateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersAwsValidateOK) GetPayload() *models.V1SpectroClusterValidatorResponse {
	return o.Payload
}

func (o *V1SpectroClustersAwsValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterValidatorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

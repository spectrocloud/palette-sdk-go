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

// V1SpectroClustersK8CertificateReader is a Reader for the V1SpectroClustersK8Certificate structure.
type V1SpectroClustersK8CertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersK8CertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersK8CertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersK8CertificateOK creates a V1SpectroClustersK8CertificateOK with default headers values
func NewV1SpectroClustersK8CertificateOK() *V1SpectroClustersK8CertificateOK {
	return &V1SpectroClustersK8CertificateOK{}
}

/*
V1SpectroClustersK8CertificateOK handles this case with default header values.

OK
*/
type V1SpectroClustersK8CertificateOK struct {
	Payload *models.V1MachineCertificates
}

func (o *V1SpectroClustersK8CertificateOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/k8certificates][%d] v1SpectroClustersK8CertificateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersK8CertificateOK) GetPayload() *models.V1MachineCertificates {
	return o.Payload
}

func (o *V1SpectroClustersK8CertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineCertificates)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
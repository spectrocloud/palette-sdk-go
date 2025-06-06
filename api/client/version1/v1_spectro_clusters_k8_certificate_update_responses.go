// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersK8CertificateUpdateReader is a Reader for the V1SpectroClustersK8CertificateUpdate structure.
type V1SpectroClustersK8CertificateUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersK8CertificateUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersK8CertificateUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersK8CertificateUpdateNoContent creates a V1SpectroClustersK8CertificateUpdateNoContent with default headers values
func NewV1SpectroClustersK8CertificateUpdateNoContent() *V1SpectroClustersK8CertificateUpdateNoContent {
	return &V1SpectroClustersK8CertificateUpdateNoContent{}
}

/*
V1SpectroClustersK8CertificateUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersK8CertificateUpdateNoContent struct {
}

func (o *V1SpectroClustersK8CertificateUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/k8certificates][%d] v1SpectroClustersK8CertificateUpdateNoContent ", 204)
}

func (o *V1SpectroClustersK8CertificateUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

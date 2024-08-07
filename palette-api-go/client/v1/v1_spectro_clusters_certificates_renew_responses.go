// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersCertificatesRenewReader is a Reader for the V1SpectroClustersCertificatesRenew structure.
type V1SpectroClustersCertificatesRenewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersCertificatesRenewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersCertificatesRenewNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersCertificatesRenewNoContent creates a V1SpectroClustersCertificatesRenewNoContent with default headers values
func NewV1SpectroClustersCertificatesRenewNoContent() *V1SpectroClustersCertificatesRenewNoContent {
	return &V1SpectroClustersCertificatesRenewNoContent{}
}

/*V1SpectroClustersCertificatesRenewNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1SpectroClustersCertificatesRenewNoContent struct {
}

func (o *V1SpectroClustersCertificatesRenewNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/spectroclusters/{uid}/k8certificates/renew][%d] v1SpectroClustersCertificatesRenewNoContent ", 204)
}

func (o *V1SpectroClustersCertificatesRenewNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

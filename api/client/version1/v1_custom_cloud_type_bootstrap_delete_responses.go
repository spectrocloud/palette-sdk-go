// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CustomCloudTypeBootstrapDeleteReader is a Reader for the V1CustomCloudTypeBootstrapDelete structure.
type V1CustomCloudTypeBootstrapDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CustomCloudTypeBootstrapDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CustomCloudTypeBootstrapDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CustomCloudTypeBootstrapDeleteNoContent creates a V1CustomCloudTypeBootstrapDeleteNoContent with default headers values
func NewV1CustomCloudTypeBootstrapDeleteNoContent() *V1CustomCloudTypeBootstrapDeleteNoContent {
	return &V1CustomCloudTypeBootstrapDeleteNoContent{}
}

/*
V1CustomCloudTypeBootstrapDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1CustomCloudTypeBootstrapDeleteNoContent struct {
}

func (o *V1CustomCloudTypeBootstrapDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/clouds/cloudTypes/{cloudType}/content/bootstrap][%d] v1CustomCloudTypeBootstrapDeleteNoContent ", 204)
}

func (o *V1CustomCloudTypeBootstrapDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

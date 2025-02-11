// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1MetricsUIDDeleteReader is a Reader for the V1MetricsUIDDelete structure.
type V1MetricsUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1MetricsUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1MetricsUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1MetricsUIDDeleteNoContent creates a V1MetricsUIDDeleteNoContent with default headers values
func NewV1MetricsUIDDeleteNoContent() *V1MetricsUIDDeleteNoContent {
	return &V1MetricsUIDDeleteNoContent{}
}

/*V1MetricsUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1MetricsUIDDeleteNoContent struct {
}

func (o *V1MetricsUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/metrics/{resourceKind}/{resourceUid}/values][%d] v1MetricsUidDeleteNoContent ", 204)
}

func (o *V1MetricsUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

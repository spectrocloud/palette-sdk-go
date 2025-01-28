// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1OverlordsUIDVsphereAccountUpdateReader is a Reader for the V1OverlordsUIDVsphereAccountUpdate structure.
type V1OverlordsUIDVsphereAccountUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OverlordsUIDVsphereAccountUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1OverlordsUIDVsphereAccountUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OverlordsUIDVsphereAccountUpdateNoContent creates a V1OverlordsUIDVsphereAccountUpdateNoContent with default headers values
func NewV1OverlordsUIDVsphereAccountUpdateNoContent() *V1OverlordsUIDVsphereAccountUpdateNoContent {
	return &V1OverlordsUIDVsphereAccountUpdateNoContent{}
}

/*V1OverlordsUIDVsphereAccountUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1OverlordsUIDVsphereAccountUpdateNoContent struct {
}

func (o *V1OverlordsUIDVsphereAccountUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/overlords/vsphere/{uid}/account][%d] v1OverlordsUidVsphereAccountUpdateNoContent ", 204)
}

func (o *V1OverlordsUIDVsphereAccountUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

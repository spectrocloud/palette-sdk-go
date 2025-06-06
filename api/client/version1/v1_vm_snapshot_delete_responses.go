// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1VMSnapshotDeleteReader is a Reader for the V1VMSnapshotDelete structure.
type V1VMSnapshotDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1VMSnapshotDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1VMSnapshotDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1VMSnapshotDeleteNoContent creates a V1VMSnapshotDeleteNoContent with default headers values
func NewV1VMSnapshotDeleteNoContent() *V1VMSnapshotDeleteNoContent {
	return &V1VMSnapshotDeleteNoContent{}
}

/*
V1VMSnapshotDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1VMSnapshotDeleteNoContent struct {
}

func (o *V1VMSnapshotDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/spectroclusters/{uid}/vms/{vmName}/snapshot/{snapshotName}][%d] v1VmSnapshotDeleteNoContent ", 204)
}

func (o *V1VMSnapshotDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

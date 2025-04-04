// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1WorkspacesUIDClusterRbacDeleteReader is a Reader for the V1WorkspacesUIDClusterRbacDelete structure.
type V1WorkspacesUIDClusterRbacDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1WorkspacesUIDClusterRbacDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1WorkspacesUIDClusterRbacDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1WorkspacesUIDClusterRbacDeleteNoContent creates a V1WorkspacesUIDClusterRbacDeleteNoContent with default headers values
func NewV1WorkspacesUIDClusterRbacDeleteNoContent() *V1WorkspacesUIDClusterRbacDeleteNoContent {
	return &V1WorkspacesUIDClusterRbacDeleteNoContent{}
}

/*
V1WorkspacesUIDClusterRbacDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1WorkspacesUIDClusterRbacDeleteNoContent struct {
}

func (o *V1WorkspacesUIDClusterRbacDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/workspaces/{uid}/clusterRbacs/{clusterRbacUid}][%d] v1WorkspacesUidClusterRbacDeleteNoContent ", 204)
}

func (o *V1WorkspacesUIDClusterRbacDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

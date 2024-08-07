// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1ClusterGroupsUIDDeleteReader is a Reader for the V1ClusterGroupsUIDDelete structure.
type V1ClusterGroupsUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterGroupsUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1ClusterGroupsUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterGroupsUIDDeleteNoContent creates a V1ClusterGroupsUIDDeleteNoContent with default headers values
func NewV1ClusterGroupsUIDDeleteNoContent() *V1ClusterGroupsUIDDeleteNoContent {
	return &V1ClusterGroupsUIDDeleteNoContent{}
}

/*
V1ClusterGroupsUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1ClusterGroupsUIDDeleteNoContent struct {
}

func (o *V1ClusterGroupsUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/clustergroups/{uid}][%d] v1ClusterGroupsUidDeleteNoContent ", 204)
}

func (o *V1ClusterGroupsUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

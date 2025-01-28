// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1ClusterGroupsUIDHostClusterUpdateReader is a Reader for the V1ClusterGroupsUIDHostClusterUpdate structure.
type V1ClusterGroupsUIDHostClusterUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterGroupsUIDHostClusterUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1ClusterGroupsUIDHostClusterUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterGroupsUIDHostClusterUpdateNoContent creates a V1ClusterGroupsUIDHostClusterUpdateNoContent with default headers values
func NewV1ClusterGroupsUIDHostClusterUpdateNoContent() *V1ClusterGroupsUIDHostClusterUpdateNoContent {
	return &V1ClusterGroupsUIDHostClusterUpdateNoContent{}
}

/*V1ClusterGroupsUIDHostClusterUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1ClusterGroupsUIDHostClusterUpdateNoContent struct {
}

func (o *V1ClusterGroupsUIDHostClusterUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/clustergroups/{uid}/hostCluster][%d] v1ClusterGroupsUidHostClusterUpdateNoContent ", 204)
}

func (o *V1ClusterGroupsUIDHostClusterUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TenantPrefClusterGroupUpdateReader is a Reader for the V1TenantPrefClusterGroupUpdate structure.
type V1TenantPrefClusterGroupUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantPrefClusterGroupUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TenantPrefClusterGroupUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantPrefClusterGroupUpdateNoContent creates a V1TenantPrefClusterGroupUpdateNoContent with default headers values
func NewV1TenantPrefClusterGroupUpdateNoContent() *V1TenantPrefClusterGroupUpdateNoContent {
	return &V1TenantPrefClusterGroupUpdateNoContent{}
}

/*
V1TenantPrefClusterGroupUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TenantPrefClusterGroupUpdateNoContent struct {
}

func (o *V1TenantPrefClusterGroupUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/tenants/{tenantUid}/preferences/clusterGroup][%d] v1TenantPrefClusterGroupUpdateNoContent ", 204)
}

func (o *V1TenantPrefClusterGroupUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

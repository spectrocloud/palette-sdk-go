// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TenantUIDAssetsCertsUIDDeleteReader is a Reader for the V1TenantUIDAssetsCertsUIDDelete structure.
type V1TenantUIDAssetsCertsUIDDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantUIDAssetsCertsUIDDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TenantUIDAssetsCertsUIDDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantUIDAssetsCertsUIDDeleteNoContent creates a V1TenantUIDAssetsCertsUIDDeleteNoContent with default headers values
func NewV1TenantUIDAssetsCertsUIDDeleteNoContent() *V1TenantUIDAssetsCertsUIDDeleteNoContent {
	return &V1TenantUIDAssetsCertsUIDDeleteNoContent{}
}

/*
V1TenantUIDAssetsCertsUIDDeleteNoContent handles this case with default header values.

The resource was deleted successfully
*/
type V1TenantUIDAssetsCertsUIDDeleteNoContent struct {
}

func (o *V1TenantUIDAssetsCertsUIDDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/tenants/{tenantUid}/assets/certs/{certificateUid}][%d] v1TenantUidAssetsCertsUidDeleteNoContent ", 204)
}

func (o *V1TenantUIDAssetsCertsUIDDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TenantUIDAssetsDataSinksUpdateReader is a Reader for the V1TenantUIDAssetsDataSinksUpdate structure.
type V1TenantUIDAssetsDataSinksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantUIDAssetsDataSinksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TenantUIDAssetsDataSinksUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantUIDAssetsDataSinksUpdateNoContent creates a V1TenantUIDAssetsDataSinksUpdateNoContent with default headers values
func NewV1TenantUIDAssetsDataSinksUpdateNoContent() *V1TenantUIDAssetsDataSinksUpdateNoContent {
	return &V1TenantUIDAssetsDataSinksUpdateNoContent{}
}

/*
V1TenantUIDAssetsDataSinksUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TenantUIDAssetsDataSinksUpdateNoContent struct {
}

func (o *V1TenantUIDAssetsDataSinksUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/tenants/{tenantUid}/assets/dataSinks][%d] v1TenantUidAssetsDataSinksUpdateNoContent ", 204)
}

func (o *V1TenantUIDAssetsDataSinksUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

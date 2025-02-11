// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1AppDeploymentsProfileTiersUIDUpdateReader is a Reader for the V1AppDeploymentsProfileTiersUIDUpdate structure.
type V1AppDeploymentsProfileTiersUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppDeploymentsProfileTiersUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1AppDeploymentsProfileTiersUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppDeploymentsProfileTiersUIDUpdateNoContent creates a V1AppDeploymentsProfileTiersUIDUpdateNoContent with default headers values
func NewV1AppDeploymentsProfileTiersUIDUpdateNoContent() *V1AppDeploymentsProfileTiersUIDUpdateNoContent {
	return &V1AppDeploymentsProfileTiersUIDUpdateNoContent{}
}

/*V1AppDeploymentsProfileTiersUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1AppDeploymentsProfileTiersUIDUpdateNoContent struct {
}

func (o *V1AppDeploymentsProfileTiersUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/appDeployments/{uid}/profile/tiers/{tierUid}][%d] v1AppDeploymentsProfileTiersUidUpdateNoContent ", 204)
}

func (o *V1AppDeploymentsProfileTiersUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

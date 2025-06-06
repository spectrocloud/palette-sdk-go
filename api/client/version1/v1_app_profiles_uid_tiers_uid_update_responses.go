// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1AppProfilesUIDTiersUIDUpdateReader is a Reader for the V1AppProfilesUIDTiersUIDUpdate structure.
type V1AppProfilesUIDTiersUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppProfilesUIDTiersUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1AppProfilesUIDTiersUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppProfilesUIDTiersUIDUpdateNoContent creates a V1AppProfilesUIDTiersUIDUpdateNoContent with default headers values
func NewV1AppProfilesUIDTiersUIDUpdateNoContent() *V1AppProfilesUIDTiersUIDUpdateNoContent {
	return &V1AppProfilesUIDTiersUIDUpdateNoContent{}
}

/*
V1AppProfilesUIDTiersUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1AppProfilesUIDTiersUIDUpdateNoContent struct {
}

func (o *V1AppProfilesUIDTiersUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/appProfiles/{uid}/tiers/{tierUid}][%d] v1AppProfilesUidTiersUidUpdateNoContent ", 204)
}

func (o *V1AppProfilesUIDTiersUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

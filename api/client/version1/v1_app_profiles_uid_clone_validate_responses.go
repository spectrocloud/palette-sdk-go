// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1AppProfilesUIDCloneValidateReader is a Reader for the V1AppProfilesUIDCloneValidate structure.
type V1AppProfilesUIDCloneValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppProfilesUIDCloneValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1AppProfilesUIDCloneValidateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppProfilesUIDCloneValidateNoContent creates a V1AppProfilesUIDCloneValidateNoContent with default headers values
func NewV1AppProfilesUIDCloneValidateNoContent() *V1AppProfilesUIDCloneValidateNoContent {
	return &V1AppProfilesUIDCloneValidateNoContent{}
}

/*
V1AppProfilesUIDCloneValidateNoContent handles this case with default header values.

Ok response without content
*/
type V1AppProfilesUIDCloneValidateNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1AppProfilesUIDCloneValidateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/appProfiles/{uid}/clone/validate][%d] v1AppProfilesUidCloneValidateNoContent ", 204)
}

func (o *V1AppProfilesUIDCloneValidateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}

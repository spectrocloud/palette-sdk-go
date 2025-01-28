// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1TeamsUIDPatchReader is a Reader for the V1TeamsUIDPatch structure.
type V1TeamsUIDPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TeamsUIDPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1TeamsUIDPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TeamsUIDPatchNoContent creates a V1TeamsUIDPatchNoContent with default headers values
func NewV1TeamsUIDPatchNoContent() *V1TeamsUIDPatchNoContent {
	return &V1TeamsUIDPatchNoContent{}
}

/*V1TeamsUIDPatchNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1TeamsUIDPatchNoContent struct {
}

func (o *V1TeamsUIDPatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/teams/{uid}][%d] v1TeamsUidPatchNoContent ", 204)
}

func (o *V1TeamsUIDPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

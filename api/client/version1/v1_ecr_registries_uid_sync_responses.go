// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1EcrRegistriesUIDSyncReader is a Reader for the V1EcrRegistriesUIDSync structure.
type V1EcrRegistriesUIDSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EcrRegistriesUIDSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewV1EcrRegistriesUIDSyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EcrRegistriesUIDSyncAccepted creates a V1EcrRegistriesUIDSyncAccepted with default headers values
func NewV1EcrRegistriesUIDSyncAccepted() *V1EcrRegistriesUIDSyncAccepted {
	return &V1EcrRegistriesUIDSyncAccepted{}
}

/*
V1EcrRegistriesUIDSyncAccepted handles this case with default header values.

Ok response without content
*/
type V1EcrRegistriesUIDSyncAccepted struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1EcrRegistriesUIDSyncAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/registries/oci/{uid}/ecr/sync][%d] v1EcrRegistriesUidSyncAccepted ", 202)
}

func (o *V1EcrRegistriesUIDSyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}

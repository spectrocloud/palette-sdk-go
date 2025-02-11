// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1BasicOciRegistriesUIDSyncReader is a Reader for the V1BasicOciRegistriesUIDSync structure.
type V1BasicOciRegistriesUIDSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1BasicOciRegistriesUIDSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewV1BasicOciRegistriesUIDSyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1BasicOciRegistriesUIDSyncAccepted creates a V1BasicOciRegistriesUIDSyncAccepted with default headers values
func NewV1BasicOciRegistriesUIDSyncAccepted() *V1BasicOciRegistriesUIDSyncAccepted {
	return &V1BasicOciRegistriesUIDSyncAccepted{}
}

/*
V1BasicOciRegistriesUIDSyncAccepted handles this case with default header values.

Ok response without content
*/
type V1BasicOciRegistriesUIDSyncAccepted struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1BasicOciRegistriesUIDSyncAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/registries/oci/{uid}/basic/sync][%d] v1BasicOciRegistriesUidSyncAccepted ", 202)
}

func (o *V1BasicOciRegistriesUIDSyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}

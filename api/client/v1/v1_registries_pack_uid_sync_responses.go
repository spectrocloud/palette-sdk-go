// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1RegistriesPackUIDSyncReader is a Reader for the V1RegistriesPackUIDSync structure.
type V1RegistriesPackUIDSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1RegistriesPackUIDSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewV1RegistriesPackUIDSyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1RegistriesPackUIDSyncAccepted creates a V1RegistriesPackUIDSyncAccepted with default headers values
func NewV1RegistriesPackUIDSyncAccepted() *V1RegistriesPackUIDSyncAccepted {
	return &V1RegistriesPackUIDSyncAccepted{}
}

/*
V1RegistriesPackUIDSyncAccepted handles this case with default header values.

Ok response without content
*/
type V1RegistriesPackUIDSyncAccepted struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1RegistriesPackUIDSyncAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/registries/pack/{uid}/sync][%d] v1RegistriesPackUidSyncAccepted ", 202)
}

func (o *V1RegistriesPackUIDSyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}

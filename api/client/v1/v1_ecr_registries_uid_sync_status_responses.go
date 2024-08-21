// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1EcrRegistriesUIDSyncStatusReader is a Reader for the V1EcrRegistriesUIDSyncStatus structure.
type V1EcrRegistriesUIDSyncStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EcrRegistriesUIDSyncStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1EcrRegistriesUIDSyncStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EcrRegistriesUIDSyncStatusOK creates a V1EcrRegistriesUIDSyncStatusOK with default headers values
func NewV1EcrRegistriesUIDSyncStatusOK() *V1EcrRegistriesUIDSyncStatusOK {
	return &V1EcrRegistriesUIDSyncStatusOK{}
}

/*
V1EcrRegistriesUIDSyncStatusOK handles this case with default header values.

Ecr registry sync status
*/
type V1EcrRegistriesUIDSyncStatusOK struct {
	Payload *models.V1RegistrySyncStatus
}

func (o *V1EcrRegistriesUIDSyncStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/registries/oci/{uid}/ecr/sync/status][%d] v1EcrRegistriesUidSyncStatusOK  %+v", 200, o.Payload)
}

func (o *V1EcrRegistriesUIDSyncStatusOK) GetPayload() *models.V1RegistrySyncStatus {
	return o.Payload
}

func (o *V1EcrRegistriesUIDSyncStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1RegistrySyncStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
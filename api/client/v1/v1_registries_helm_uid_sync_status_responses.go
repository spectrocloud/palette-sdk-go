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

// V1RegistriesHelmUIDSyncStatusReader is a Reader for the V1RegistriesHelmUIDSyncStatus structure.
type V1RegistriesHelmUIDSyncStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1RegistriesHelmUIDSyncStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1RegistriesHelmUIDSyncStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1RegistriesHelmUIDSyncStatusOK creates a V1RegistriesHelmUIDSyncStatusOK with default headers values
func NewV1RegistriesHelmUIDSyncStatusOK() *V1RegistriesHelmUIDSyncStatusOK {
	return &V1RegistriesHelmUIDSyncStatusOK{}
}

/*
V1RegistriesHelmUIDSyncStatusOK handles this case with default header values.

Helm registry sync status
*/
type V1RegistriesHelmUIDSyncStatusOK struct {
	Payload *models.V1RegistrySyncStatus
}

func (o *V1RegistriesHelmUIDSyncStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/registries/helm/{uid}/sync/status][%d] v1RegistriesHelmUidSyncStatusOK  %+v", 200, o.Payload)
}

func (o *V1RegistriesHelmUIDSyncStatusOK) GetPayload() *models.V1RegistrySyncStatus {
	return o.Payload
}

func (o *V1RegistriesHelmUIDSyncStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1RegistrySyncStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// V1BasicOciRegistriesUIDSyncStatusReader is a Reader for the V1BasicOciRegistriesUIDSyncStatus structure.
type V1BasicOciRegistriesUIDSyncStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1BasicOciRegistriesUIDSyncStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1BasicOciRegistriesUIDSyncStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1BasicOciRegistriesUIDSyncStatusOK creates a V1BasicOciRegistriesUIDSyncStatusOK with default headers values
func NewV1BasicOciRegistriesUIDSyncStatusOK() *V1BasicOciRegistriesUIDSyncStatusOK {
	return &V1BasicOciRegistriesUIDSyncStatusOK{}
}

/*V1BasicOciRegistriesUIDSyncStatusOK handles this case with default header values.

Oci registry sync status
*/
type V1BasicOciRegistriesUIDSyncStatusOK struct {
	Payload *models.V1RegistrySyncStatus
}

func (o *V1BasicOciRegistriesUIDSyncStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/registries/oci/{uid}/basic/sync/status][%d] v1BasicOciRegistriesUidSyncStatusOK  %+v", 200, o.Payload)
}

func (o *V1BasicOciRegistriesUIDSyncStatusOK) GetPayload() *models.V1RegistrySyncStatus {
	return o.Payload
}

func (o *V1BasicOciRegistriesUIDSyncStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1RegistrySyncStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

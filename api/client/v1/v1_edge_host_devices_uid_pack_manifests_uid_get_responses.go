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

// V1EdgeHostDevicesUIDPackManifestsUIDGetReader is a Reader for the V1EdgeHostDevicesUIDPackManifestsUIDGet structure.
type V1EdgeHostDevicesUIDPackManifestsUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EdgeHostDevicesUIDPackManifestsUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1EdgeHostDevicesUIDPackManifestsUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EdgeHostDevicesUIDPackManifestsUIDGetOK creates a V1EdgeHostDevicesUIDPackManifestsUIDGetOK with default headers values
func NewV1EdgeHostDevicesUIDPackManifestsUIDGetOK() *V1EdgeHostDevicesUIDPackManifestsUIDGetOK {
	return &V1EdgeHostDevicesUIDPackManifestsUIDGetOK{}
}

/*V1EdgeHostDevicesUIDPackManifestsUIDGetOK handles this case with default header values.

Pack manifest content
*/
type V1EdgeHostDevicesUIDPackManifestsUIDGetOK struct {
	Payload *models.V1Manifest
}

func (o *V1EdgeHostDevicesUIDPackManifestsUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/edgehosts/{uid}/pack/manifests/{manifestUid}][%d] v1EdgeHostDevicesUidPackManifestsUidGetOK  %+v", 200, o.Payload)
}

func (o *V1EdgeHostDevicesUIDPackManifestsUIDGetOK) GetPayload() *models.V1Manifest {
	return o.Payload
}

func (o *V1EdgeHostDevicesUIDPackManifestsUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

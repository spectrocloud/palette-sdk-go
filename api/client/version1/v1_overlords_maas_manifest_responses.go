// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1OverlordsMaasManifestReader is a Reader for the V1OverlordsMaasManifest structure.
type V1OverlordsMaasManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OverlordsMaasManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OverlordsMaasManifestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OverlordsMaasManifestOK creates a V1OverlordsMaasManifestOK with default headers values
func NewV1OverlordsMaasManifestOK() *V1OverlordsMaasManifestOK {
	return &V1OverlordsMaasManifestOK{}
}

/*
V1OverlordsMaasManifestOK handles this case with default header values.

(empty)
*/
type V1OverlordsMaasManifestOK struct {
	Payload *models.V1OverlordManifest
}

func (o *V1OverlordsMaasManifestOK) Error() string {
	return fmt.Sprintf("[GET /v1/overlords/maas/manifest][%d] v1OverlordsMaasManifestOK  %+v", 200, o.Payload)
}

func (o *V1OverlordsMaasManifestOK) GetPayload() *models.V1OverlordManifest {
	return o.Payload
}

func (o *V1OverlordsMaasManifestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OverlordManifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

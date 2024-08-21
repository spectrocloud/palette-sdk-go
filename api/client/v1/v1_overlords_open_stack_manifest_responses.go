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

// V1OverlordsOpenStackManifestReader is a Reader for the V1OverlordsOpenStackManifest structure.
type V1OverlordsOpenStackManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OverlordsOpenStackManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OverlordsOpenStackManifestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OverlordsOpenStackManifestOK creates a V1OverlordsOpenStackManifestOK with default headers values
func NewV1OverlordsOpenStackManifestOK() *V1OverlordsOpenStackManifestOK {
	return &V1OverlordsOpenStackManifestOK{}
}

/*
V1OverlordsOpenStackManifestOK handles this case with default header values.

(empty)
*/
type V1OverlordsOpenStackManifestOK struct {
	Payload *models.V1OverlordManifest
}

func (o *V1OverlordsOpenStackManifestOK) Error() string {
	return fmt.Sprintf("[GET /v1/overlords/openstack/manifest][%d] v1OverlordsOpenStackManifestOK  %+v", 200, o.Payload)
}

func (o *V1OverlordsOpenStackManifestOK) GetPayload() *models.V1OverlordManifest {
	return o.Payload
}

func (o *V1OverlordsOpenStackManifestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1OverlordManifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// V1SpectroClustersUIDPackManifestsUIDGetReader is a Reader for the V1SpectroClustersUIDPackManifestsUIDGet structure.
type V1SpectroClustersUIDPackManifestsUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDPackManifestsUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersUIDPackManifestsUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDPackManifestsUIDGetOK creates a V1SpectroClustersUIDPackManifestsUIDGetOK with default headers values
func NewV1SpectroClustersUIDPackManifestsUIDGetOK() *V1SpectroClustersUIDPackManifestsUIDGetOK {
	return &V1SpectroClustersUIDPackManifestsUIDGetOK{}
}

/*
V1SpectroClustersUIDPackManifestsUIDGetOK handles this case with default header values.

Pack manifest content
*/
type V1SpectroClustersUIDPackManifestsUIDGetOK struct {
	Payload *models.V1Manifest
}

func (o *V1SpectroClustersUIDPackManifestsUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/pack/manifests/{manifestUid}][%d] v1SpectroClustersUidPackManifestsUidGetOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersUIDPackManifestsUIDGetOK) GetPayload() *models.V1Manifest {
	return o.Payload
}

func (o *V1SpectroClustersUIDPackManifestsUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

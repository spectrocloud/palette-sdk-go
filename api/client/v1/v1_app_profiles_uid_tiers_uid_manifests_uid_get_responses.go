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

// V1AppProfilesUIDTiersUIDManifestsUIDGetReader is a Reader for the V1AppProfilesUIDTiersUIDManifestsUIDGet structure.
type V1AppProfilesUIDTiersUIDManifestsUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AppProfilesUIDTiersUIDManifestsUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsUIDGetOK creates a V1AppProfilesUIDTiersUIDManifestsUIDGetOK with default headers values
func NewV1AppProfilesUIDTiersUIDManifestsUIDGetOK() *V1AppProfilesUIDTiersUIDManifestsUIDGetOK {
	return &V1AppProfilesUIDTiersUIDManifestsUIDGetOK{}
}

/*
V1AppProfilesUIDTiersUIDManifestsUIDGetOK handles this case with default header values.

OK
*/
type V1AppProfilesUIDTiersUIDManifestsUIDGetOK struct {
	Payload *models.V1Manifest
}

func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/appProfiles/{uid}/tiers/{tierUid}/manifests/{manifestUid}][%d] v1AppProfilesUidTiersUidManifestsUidGetOK  %+v", 200, o.Payload)
}

func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetOK) GetPayload() *models.V1Manifest {
	return o.Payload
}

func (o *V1AppProfilesUIDTiersUIDManifestsUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// V1AppProfilesUIDTiersUIDManifestsGetReader is a Reader for the V1AppProfilesUIDTiersUIDManifestsGet structure.
type V1AppProfilesUIDTiersUIDManifestsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppProfilesUIDTiersUIDManifestsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AppProfilesUIDTiersUIDManifestsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppProfilesUIDTiersUIDManifestsGetOK creates a V1AppProfilesUIDTiersUIDManifestsGetOK with default headers values
func NewV1AppProfilesUIDTiersUIDManifestsGetOK() *V1AppProfilesUIDTiersUIDManifestsGetOK {
	return &V1AppProfilesUIDTiersUIDManifestsGetOK{}
}

/*
V1AppProfilesUIDTiersUIDManifestsGetOK handles this case with default header values.

OK
*/
type V1AppProfilesUIDTiersUIDManifestsGetOK struct {
	Payload *models.V1AppTierManifests
}

func (o *V1AppProfilesUIDTiersUIDManifestsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/appProfiles/{uid}/tiers/{tierUid}/manifests][%d] v1AppProfilesUidTiersUidManifestsGetOK  %+v", 200, o.Payload)
}

func (o *V1AppProfilesUIDTiersUIDManifestsGetOK) GetPayload() *models.V1AppTierManifests {
	return o.Payload
}

func (o *V1AppProfilesUIDTiersUIDManifestsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AppTierManifests)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

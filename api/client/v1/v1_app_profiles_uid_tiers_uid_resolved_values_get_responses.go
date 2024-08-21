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

// V1AppProfilesUIDTiersUIDResolvedValuesGetReader is a Reader for the V1AppProfilesUIDTiersUIDResolvedValuesGet structure.
type V1AppProfilesUIDTiersUIDResolvedValuesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppProfilesUIDTiersUIDResolvedValuesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AppProfilesUIDTiersUIDResolvedValuesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppProfilesUIDTiersUIDResolvedValuesGetOK creates a V1AppProfilesUIDTiersUIDResolvedValuesGetOK with default headers values
func NewV1AppProfilesUIDTiersUIDResolvedValuesGetOK() *V1AppProfilesUIDTiersUIDResolvedValuesGetOK {
	return &V1AppProfilesUIDTiersUIDResolvedValuesGetOK{}
}

/*
V1AppProfilesUIDTiersUIDResolvedValuesGetOK handles this case with default header values.

OK
*/
type V1AppProfilesUIDTiersUIDResolvedValuesGetOK struct {
	Payload *models.V1AppTierResolvedValues
}

func (o *V1AppProfilesUIDTiersUIDResolvedValuesGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/appProfiles/{uid}/tiers/{tierUid}/resolvedValues][%d] v1AppProfilesUidTiersUidResolvedValuesGetOK  %+v", 200, o.Payload)
}

func (o *V1AppProfilesUIDTiersUIDResolvedValuesGetOK) GetPayload() *models.V1AppTierResolvedValues {
	return o.Payload
}

func (o *V1AppProfilesUIDTiersUIDResolvedValuesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AppTierResolvedValues)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
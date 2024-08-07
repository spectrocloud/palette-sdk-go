// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1AppProfilesUIDTiersGetReader is a Reader for the V1AppProfilesUIDTiersGet structure.
type V1AppProfilesUIDTiersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppProfilesUIDTiersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AppProfilesUIDTiersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppProfilesUIDTiersGetOK creates a V1AppProfilesUIDTiersGetOK with default headers values
func NewV1AppProfilesUIDTiersGetOK() *V1AppProfilesUIDTiersGetOK {
	return &V1AppProfilesUIDTiersGetOK{}
}

/*V1AppProfilesUIDTiersGetOK handles this case with default header values.

OK
*/
type V1AppProfilesUIDTiersGetOK struct {
	Payload *models.V1AppProfileTiers
}

func (o *V1AppProfilesUIDTiersGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/appProfiles/{uid}/tiers][%d] v1AppProfilesUidTiersGetOK  %+v", 200, o.Payload)
}

func (o *V1AppProfilesUIDTiersGetOK) GetPayload() *models.V1AppProfileTiers {
	return o.Payload
}

func (o *V1AppProfilesUIDTiersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AppProfileTiers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

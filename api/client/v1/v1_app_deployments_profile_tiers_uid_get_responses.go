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

// V1AppDeploymentsProfileTiersUIDGetReader is a Reader for the V1AppDeploymentsProfileTiersUIDGet structure.
type V1AppDeploymentsProfileTiersUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppDeploymentsProfileTiersUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AppDeploymentsProfileTiersUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppDeploymentsProfileTiersUIDGetOK creates a V1AppDeploymentsProfileTiersUIDGetOK with default headers values
func NewV1AppDeploymentsProfileTiersUIDGetOK() *V1AppDeploymentsProfileTiersUIDGetOK {
	return &V1AppDeploymentsProfileTiersUIDGetOK{}
}

/*V1AppDeploymentsProfileTiersUIDGetOK handles this case with default header values.

OK
*/
type V1AppDeploymentsProfileTiersUIDGetOK struct {
	Payload *models.V1AppTier
}

func (o *V1AppDeploymentsProfileTiersUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/appDeployments/{uid}/profile/tiers/{tierUid}][%d] v1AppDeploymentsProfileTiersUidGetOK  %+v", 200, o.Payload)
}

func (o *V1AppDeploymentsProfileTiersUIDGetOK) GetPayload() *models.V1AppTier {
	return o.Payload
}

func (o *V1AppDeploymentsProfileTiersUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AppTier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

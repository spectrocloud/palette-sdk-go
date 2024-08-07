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

// V1AppDeploymentsProfileTiersUIDManifestsGetReader is a Reader for the V1AppDeploymentsProfileTiersUIDManifestsGet structure.
type V1AppDeploymentsProfileTiersUIDManifestsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppDeploymentsProfileTiersUIDManifestsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AppDeploymentsProfileTiersUIDManifestsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppDeploymentsProfileTiersUIDManifestsGetOK creates a V1AppDeploymentsProfileTiersUIDManifestsGetOK with default headers values
func NewV1AppDeploymentsProfileTiersUIDManifestsGetOK() *V1AppDeploymentsProfileTiersUIDManifestsGetOK {
	return &V1AppDeploymentsProfileTiersUIDManifestsGetOK{}
}

/*
V1AppDeploymentsProfileTiersUIDManifestsGetOK handles this case with default header values.

OK
*/
type V1AppDeploymentsProfileTiersUIDManifestsGetOK struct {
	Payload *models.V1AppTierManifests
}

func (o *V1AppDeploymentsProfileTiersUIDManifestsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/appDeployments/{uid}/profile/tiers/{tierUid}/manifests][%d] v1AppDeploymentsProfileTiersUidManifestsGetOK  %+v", 200, o.Payload)
}

func (o *V1AppDeploymentsProfileTiersUIDManifestsGetOK) GetPayload() *models.V1AppTierManifests {
	return o.Payload
}

func (o *V1AppDeploymentsProfileTiersUIDManifestsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AppTierManifests)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

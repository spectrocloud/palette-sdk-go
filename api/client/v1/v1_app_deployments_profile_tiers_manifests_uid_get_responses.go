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

// V1AppDeploymentsProfileTiersManifestsUIDGetReader is a Reader for the V1AppDeploymentsProfileTiersManifestsUIDGet structure.
type V1AppDeploymentsProfileTiersManifestsUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppDeploymentsProfileTiersManifestsUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AppDeploymentsProfileTiersManifestsUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppDeploymentsProfileTiersManifestsUIDGetOK creates a V1AppDeploymentsProfileTiersManifestsUIDGetOK with default headers values
func NewV1AppDeploymentsProfileTiersManifestsUIDGetOK() *V1AppDeploymentsProfileTiersManifestsUIDGetOK {
	return &V1AppDeploymentsProfileTiersManifestsUIDGetOK{}
}

/*
V1AppDeploymentsProfileTiersManifestsUIDGetOK handles this case with default header values.

OK
*/
type V1AppDeploymentsProfileTiersManifestsUIDGetOK struct {
	Payload *models.V1Manifest
}

func (o *V1AppDeploymentsProfileTiersManifestsUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/appDeployments/{uid}/profile/tiers/{tierUid}/manifests/{manifestUid}][%d] v1AppDeploymentsProfileTiersManifestsUidGetOK  %+v", 200, o.Payload)
}

func (o *V1AppDeploymentsProfileTiersManifestsUIDGetOK) GetPayload() *models.V1Manifest {
	return o.Payload
}

func (o *V1AppDeploymentsProfileTiersManifestsUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
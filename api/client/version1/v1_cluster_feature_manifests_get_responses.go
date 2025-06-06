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

// V1ClusterFeatureManifestsGetReader is a Reader for the V1ClusterFeatureManifestsGet structure.
type V1ClusterFeatureManifestsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterFeatureManifestsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterFeatureManifestsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterFeatureManifestsGetOK creates a V1ClusterFeatureManifestsGetOK with default headers values
func NewV1ClusterFeatureManifestsGetOK() *V1ClusterFeatureManifestsGetOK {
	return &V1ClusterFeatureManifestsGetOK{}
}

/*
V1ClusterFeatureManifestsGetOK handles this case with default header values.

OK
*/
type V1ClusterFeatureManifestsGetOK struct {
	Payload *models.V1ClusterManifests
}

func (o *V1ClusterFeatureManifestsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/features/manifests][%d] v1ClusterFeatureManifestsGetOK  %+v", 200, o.Payload)
}

func (o *V1ClusterFeatureManifestsGetOK) GetPayload() *models.V1ClusterManifests {
	return o.Payload
}

func (o *V1ClusterFeatureManifestsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterManifests)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

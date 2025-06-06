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

// V1CustomCloudTypeControlPlaneGetReader is a Reader for the V1CustomCloudTypeControlPlaneGet structure.
type V1CustomCloudTypeControlPlaneGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CustomCloudTypeControlPlaneGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CustomCloudTypeControlPlaneGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CustomCloudTypeControlPlaneGetOK creates a V1CustomCloudTypeControlPlaneGetOK with default headers values
func NewV1CustomCloudTypeControlPlaneGetOK() *V1CustomCloudTypeControlPlaneGetOK {
	return &V1CustomCloudTypeControlPlaneGetOK{}
}

/*
V1CustomCloudTypeControlPlaneGetOK handles this case with default header values.

(empty)
*/
type V1CustomCloudTypeControlPlaneGetOK struct {
	Payload *models.V1CustomCloudTypeContentResponse
}

func (o *V1CustomCloudTypeControlPlaneGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/cloudTypes/{cloudType}/content/controlPlane][%d] v1CustomCloudTypeControlPlaneGetOK  %+v", 200, o.Payload)
}

func (o *V1CustomCloudTypeControlPlaneGetOK) GetPayload() *models.V1CustomCloudTypeContentResponse {
	return o.Payload
}

func (o *V1CustomCloudTypeControlPlaneGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CustomCloudTypeContentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1CustomCloudTypeControlPlanePoolTemplateGetReader is a Reader for the V1CustomCloudTypeControlPlanePoolTemplateGet structure.
type V1CustomCloudTypeControlPlanePoolTemplateGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CustomCloudTypeControlPlanePoolTemplateGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1CustomCloudTypeControlPlanePoolTemplateGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CustomCloudTypeControlPlanePoolTemplateGetOK creates a V1CustomCloudTypeControlPlanePoolTemplateGetOK with default headers values
func NewV1CustomCloudTypeControlPlanePoolTemplateGetOK() *V1CustomCloudTypeControlPlanePoolTemplateGetOK {
	return &V1CustomCloudTypeControlPlanePoolTemplateGetOK{}
}

/*
V1CustomCloudTypeControlPlanePoolTemplateGetOK handles this case with default header values.

(empty)
*/
type V1CustomCloudTypeControlPlanePoolTemplateGetOK struct {
	Payload *models.V1CustomCloudTypeContentResponse
}

func (o *V1CustomCloudTypeControlPlanePoolTemplateGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/cloudTypes/{cloudType}/content/templates/controlPlanePoolTemplate][%d] v1CustomCloudTypeControlPlanePoolTemplateGetOK  %+v", 200, o.Payload)
}

func (o *V1CustomCloudTypeControlPlanePoolTemplateGetOK) GetPayload() *models.V1CustomCloudTypeContentResponse {
	return o.Payload
}

func (o *V1CustomCloudTypeControlPlanePoolTemplateGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CustomCloudTypeContentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

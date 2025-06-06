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

// V1ProjectClusterSettingsGetReader is a Reader for the V1ProjectClusterSettingsGet structure.
type V1ProjectClusterSettingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ProjectClusterSettingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ProjectClusterSettingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ProjectClusterSettingsGetOK creates a V1ProjectClusterSettingsGetOK with default headers values
func NewV1ProjectClusterSettingsGetOK() *V1ProjectClusterSettingsGetOK {
	return &V1ProjectClusterSettingsGetOK{}
}

/*
V1ProjectClusterSettingsGetOK handles this case with default header values.

(empty)
*/
type V1ProjectClusterSettingsGetOK struct {
	Payload *models.V1ProjectClusterSettings
}

func (o *V1ProjectClusterSettingsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{uid}/preferences/clusterSettings][%d] v1ProjectClusterSettingsGetOK  %+v", 200, o.Payload)
}

func (o *V1ProjectClusterSettingsGetOK) GetPayload() *models.V1ProjectClusterSettings {
	return o.Payload
}

func (o *V1ProjectClusterSettingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ProjectClusterSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1DashboardAppProfilesReader is a Reader for the V1DashboardAppProfiles structure.
type V1DashboardAppProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DashboardAppProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DashboardAppProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1DashboardAppProfilesOK creates a V1DashboardAppProfilesOK with default headers values
func NewV1DashboardAppProfilesOK() *V1DashboardAppProfilesOK {
	return &V1DashboardAppProfilesOK{}
}

/*
V1DashboardAppProfilesOK handles this case with default header values.

An array of application profiles summary items
*/
type V1DashboardAppProfilesOK struct {
	Payload *models.V1AppProfilesSummary
}

func (o *V1DashboardAppProfilesOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/appProfiles][%d] v1DashboardAppProfilesOK  %+v", 200, o.Payload)
}

func (o *V1DashboardAppProfilesOK) GetPayload() *models.V1AppProfilesSummary {
	return o.Payload
}

func (o *V1DashboardAppProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AppProfilesSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

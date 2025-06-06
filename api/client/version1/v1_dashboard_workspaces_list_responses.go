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

// V1DashboardWorkspacesListReader is a Reader for the V1DashboardWorkspacesList structure.
type V1DashboardWorkspacesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1DashboardWorkspacesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1DashboardWorkspacesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1DashboardWorkspacesListOK creates a V1DashboardWorkspacesListOK with default headers values
func NewV1DashboardWorkspacesListOK() *V1DashboardWorkspacesListOK {
	return &V1DashboardWorkspacesListOK{}
}

/*
V1DashboardWorkspacesListOK handles this case with default header values.

An array of workspace
*/
type V1DashboardWorkspacesListOK struct {
	Payload *models.V1DashboardWorkspaces
}

func (o *V1DashboardWorkspacesListOK) Error() string {
	return fmt.Sprintf("[GET /v1/dashboard/workspaces][%d] v1DashboardWorkspacesListOK  %+v", 200, o.Payload)
}

func (o *V1DashboardWorkspacesListOK) GetPayload() *models.V1DashboardWorkspaces {
	return o.Payload
}

func (o *V1DashboardWorkspacesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1DashboardWorkspaces)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

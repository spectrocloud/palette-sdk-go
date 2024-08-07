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

// V1GcpProjectsReader is a Reader for the V1GcpProjects structure.
type V1GcpProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1GcpProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1GcpProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1GcpProjectsOK creates a V1GcpProjectsOK with default headers values
func NewV1GcpProjectsOK() *V1GcpProjectsOK {
	return &V1GcpProjectsOK{}
}

/*V1GcpProjectsOK handles this case with default header values.

(empty)
*/
type V1GcpProjectsOK struct {
	Payload *models.V1GcpProjects
}

func (o *V1GcpProjectsOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/gcp/projects][%d] v1GcpProjectsOK  %+v", 200, o.Payload)
}

func (o *V1GcpProjectsOK) GetPayload() *models.V1GcpProjects {
	return o.Payload
}

func (o *V1GcpProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GcpProjects)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

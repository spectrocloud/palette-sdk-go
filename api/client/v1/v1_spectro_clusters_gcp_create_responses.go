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

// V1SpectroClustersGcpCreateReader is a Reader for the V1SpectroClustersGcpCreate structure.
type V1SpectroClustersGcpCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersGcpCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV1SpectroClustersGcpCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersGcpCreateCreated creates a V1SpectroClustersGcpCreateCreated with default headers values
func NewV1SpectroClustersGcpCreateCreated() *V1SpectroClustersGcpCreateCreated {
	return &V1SpectroClustersGcpCreateCreated{}
}

/*V1SpectroClustersGcpCreateCreated handles this case with default header values.

Created successfully
*/
type V1SpectroClustersGcpCreateCreated struct {
	/*Audit uid for the request
	 */
	AuditUID string

	Payload *models.V1UID
}

func (o *V1SpectroClustersGcpCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/gcp][%d] v1SpectroClustersGcpCreateCreated  %+v", 201, o.Payload)
}

func (o *V1SpectroClustersGcpCreateCreated) GetPayload() *models.V1UID {
	return o.Payload
}

func (o *V1SpectroClustersGcpCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	o.Payload = new(models.V1UID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1EventsComponentsObjTypeUIDListReader is a Reader for the V1EventsComponentsObjTypeUIDList structure.
type V1EventsComponentsObjTypeUIDListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EventsComponentsObjTypeUIDListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1EventsComponentsObjTypeUIDListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EventsComponentsObjTypeUIDListOK creates a V1EventsComponentsObjTypeUIDListOK with default headers values
func NewV1EventsComponentsObjTypeUIDListOK() *V1EventsComponentsObjTypeUIDListOK {
	return &V1EventsComponentsObjTypeUIDListOK{}
}

/*V1EventsComponentsObjTypeUIDListOK handles this case with default header values.

An array of component event items
*/
type V1EventsComponentsObjTypeUIDListOK struct {
	Payload *models.V1Events
}

func (o *V1EventsComponentsObjTypeUIDListOK) Error() string {
	return fmt.Sprintf("[GET /v1/events/components/{objectKind}/{objectUid}][%d] v1EventsComponentsObjTypeUidListOK  %+v", 200, o.Payload)
}

func (o *V1EventsComponentsObjTypeUIDListOK) GetPayload() *models.V1Events {
	return o.Payload
}

func (o *V1EventsComponentsObjTypeUIDListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Events)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

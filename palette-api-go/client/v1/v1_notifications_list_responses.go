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

// V1NotificationsListReader is a Reader for the V1NotificationsList structure.
type V1NotificationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1NotificationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1NotificationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1NotificationsListOK creates a V1NotificationsListOK with default headers values
func NewV1NotificationsListOK() *V1NotificationsListOK {
	return &V1NotificationsListOK{}
}

/*V1NotificationsListOK handles this case with default header values.

An array of notification items
*/
type V1NotificationsListOK struct {
	Payload *models.V1Notifications
}

func (o *V1NotificationsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/notifications/][%d] v1NotificationsListOK  %+v", 200, o.Payload)
}

func (o *V1NotificationsListOK) GetPayload() *models.V1Notifications {
	return o.Payload
}

func (o *V1NotificationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Notifications)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1NotificationsObjTypeUIDListReader is a Reader for the V1NotificationsObjTypeUIDList structure.
type V1NotificationsObjTypeUIDListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1NotificationsObjTypeUIDListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1NotificationsObjTypeUIDListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1NotificationsObjTypeUIDListOK creates a V1NotificationsObjTypeUIDListOK with default headers values
func NewV1NotificationsObjTypeUIDListOK() *V1NotificationsObjTypeUIDListOK {
	return &V1NotificationsObjTypeUIDListOK{}
}

/*
V1NotificationsObjTypeUIDListOK handles this case with default header values.

An array of component event items
*/
type V1NotificationsObjTypeUIDListOK struct {
	Payload *models.V1Notifications
}

func (o *V1NotificationsObjTypeUIDListOK) Error() string {
	return fmt.Sprintf("[GET /v1/notifications/{objectKind}/{objectUid}][%d] v1NotificationsObjTypeUidListOK  %+v", 200, o.Payload)
}

func (o *V1NotificationsObjTypeUIDListOK) GetPayload() *models.V1Notifications {
	return o.Payload
}

func (o *V1NotificationsObjTypeUIDListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1Notifications)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

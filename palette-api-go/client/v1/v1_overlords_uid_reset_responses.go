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

// V1OverlordsUIDResetReader is a Reader for the V1OverlordsUIDReset structure.
type V1OverlordsUIDResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1OverlordsUIDResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1OverlordsUIDResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1OverlordsUIDResetOK creates a V1OverlordsUIDResetOK with default headers values
func NewV1OverlordsUIDResetOK() *V1OverlordsUIDResetOK {
	return &V1OverlordsUIDResetOK{}
}

/*V1OverlordsUIDResetOK handles this case with default header values.

(empty)
*/
type V1OverlordsUIDResetOK struct {
	Payload *models.V1UpdatedMsg
}

func (o *V1OverlordsUIDResetOK) Error() string {
	return fmt.Sprintf("[PUT /v1/overlords/{uid}/reset][%d] v1OverlordsUidResetOK  %+v", 200, o.Payload)
}

func (o *V1OverlordsUIDResetOK) GetPayload() *models.V1UpdatedMsg {
	return o.Payload
}

func (o *V1OverlordsUIDResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UpdatedMsg)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

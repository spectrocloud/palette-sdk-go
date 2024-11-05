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

// V1UsersKubectlSessionUIDReader is a Reader for the V1UsersKubectlSessionUID structure.
type V1UsersKubectlSessionUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersKubectlSessionUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1UsersKubectlSessionUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersKubectlSessionUIDOK creates a V1UsersKubectlSessionUIDOK with default headers values
func NewV1UsersKubectlSessionUIDOK() *V1UsersKubectlSessionUIDOK {
	return &V1UsersKubectlSessionUIDOK{}
}

/*V1UsersKubectlSessionUIDOK handles this case with default header values.

OK
*/
type V1UsersKubectlSessionUIDOK struct {
	Payload *models.V1UserKubectlSession
}

func (o *V1UsersKubectlSessionUIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/kubectl/session/{sessionUid}][%d] v1UsersKubectlSessionUidOK  %+v", 200, o.Payload)
}

func (o *V1UsersKubectlSessionUIDOK) GetPayload() *models.V1UserKubectlSession {
	return o.Payload
}

func (o *V1UsersKubectlSessionUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UserKubectlSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

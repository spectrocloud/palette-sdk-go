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

// V1UsersAssetsLocationS3CreateReader is a Reader for the V1UsersAssetsLocationS3Create structure.
type V1UsersAssetsLocationS3CreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersAssetsLocationS3CreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV1UsersAssetsLocationS3CreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersAssetsLocationS3CreateCreated creates a V1UsersAssetsLocationS3CreateCreated with default headers values
func NewV1UsersAssetsLocationS3CreateCreated() *V1UsersAssetsLocationS3CreateCreated {
	return &V1UsersAssetsLocationS3CreateCreated{}
}

/*V1UsersAssetsLocationS3CreateCreated handles this case with default header values.

Created successfully
*/
type V1UsersAssetsLocationS3CreateCreated struct {
	/*Audit uid for the request
	 */
	AuditUID string

	Payload *models.V1UID
}

func (o *V1UsersAssetsLocationS3CreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/users/assets/locations/s3][%d] v1UsersAssetsLocationS3CreateCreated  %+v", 201, o.Payload)
}

func (o *V1UsersAssetsLocationS3CreateCreated) GetPayload() *models.V1UID {
	return o.Payload
}

func (o *V1UsersAssetsLocationS3CreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	o.Payload = new(models.V1UID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

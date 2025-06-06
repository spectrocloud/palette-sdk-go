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

// V1UsersAssetsLocationS3GetReader is a Reader for the V1UsersAssetsLocationS3Get structure.
type V1UsersAssetsLocationS3GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersAssetsLocationS3GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1UsersAssetsLocationS3GetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersAssetsLocationS3GetOK creates a V1UsersAssetsLocationS3GetOK with default headers values
func NewV1UsersAssetsLocationS3GetOK() *V1UsersAssetsLocationS3GetOK {
	return &V1UsersAssetsLocationS3GetOK{}
}

/*
V1UsersAssetsLocationS3GetOK handles this case with default header values.

(empty)
*/
type V1UsersAssetsLocationS3GetOK struct {
	Payload *models.V1UserAssetsLocationS3
}

func (o *V1UsersAssetsLocationS3GetOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/assets/locations/s3/{uid}][%d] v1UsersAssetsLocationS3GetOK  %+v", 200, o.Payload)
}

func (o *V1UsersAssetsLocationS3GetOK) GetPayload() *models.V1UserAssetsLocationS3 {
	return o.Payload
}

func (o *V1UsersAssetsLocationS3GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UserAssetsLocationS3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

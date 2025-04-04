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

// V1UsersAssetsLocationMinioGetReader is a Reader for the V1UsersAssetsLocationMinioGet structure.
type V1UsersAssetsLocationMinioGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersAssetsLocationMinioGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1UsersAssetsLocationMinioGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersAssetsLocationMinioGetOK creates a V1UsersAssetsLocationMinioGetOK with default headers values
func NewV1UsersAssetsLocationMinioGetOK() *V1UsersAssetsLocationMinioGetOK {
	return &V1UsersAssetsLocationMinioGetOK{}
}

/*
V1UsersAssetsLocationMinioGetOK handles this case with default header values.

(empty)
*/
type V1UsersAssetsLocationMinioGetOK struct {
	Payload *models.V1UserAssetsLocationS3
}

func (o *V1UsersAssetsLocationMinioGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/assets/locations/minio/{uid}][%d] v1UsersAssetsLocationMinioGetOK  %+v", 200, o.Payload)
}

func (o *V1UsersAssetsLocationMinioGetOK) GetPayload() *models.V1UserAssetsLocationS3 {
	return o.Payload
}

func (o *V1UsersAssetsLocationMinioGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UserAssetsLocationS3)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

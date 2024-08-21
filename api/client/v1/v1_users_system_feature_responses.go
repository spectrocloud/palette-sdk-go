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

// V1UsersSystemFeatureReader is a Reader for the V1UsersSystemFeature structure.
type V1UsersSystemFeatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1UsersSystemFeatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1UsersSystemFeatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1UsersSystemFeatureOK creates a V1UsersSystemFeatureOK with default headers values
func NewV1UsersSystemFeatureOK() *V1UsersSystemFeatureOK {
	return &V1UsersSystemFeatureOK{}
}

/*
V1UsersSystemFeatureOK handles this case with default header values.

OK
*/
type V1UsersSystemFeatureOK struct {
	Payload *models.V1SystemFeatures
}

func (o *V1UsersSystemFeatureOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/system/features][%d] v1UsersSystemFeatureOK  %+v", 200, o.Payload)
}

func (o *V1UsersSystemFeatureOK) GetPayload() *models.V1SystemFeatures {
	return o.Payload
}

func (o *V1UsersSystemFeatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SystemFeatures)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
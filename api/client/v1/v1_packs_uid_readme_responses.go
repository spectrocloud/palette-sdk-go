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

// V1PacksUIDReadmeReader is a Reader for the V1PacksUIDReadme structure.
type V1PacksUIDReadmeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1PacksUIDReadmeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1PacksUIDReadmeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1PacksUIDReadmeOK creates a V1PacksUIDReadmeOK with default headers values
func NewV1PacksUIDReadmeOK() *V1PacksUIDReadmeOK {
	return &V1PacksUIDReadmeOK{}
}

/*
V1PacksUIDReadmeOK handles this case with default header values.

Readme describes the documentation of the specified pack
*/
type V1PacksUIDReadmeOK struct {
	Payload *models.V1PackReadme
}

func (o *V1PacksUIDReadmeOK) Error() string {
	return fmt.Sprintf("[GET /v1/packs/{uid}/readme][%d] v1PacksUidReadmeOK  %+v", 200, o.Payload)
}

func (o *V1PacksUIDReadmeOK) GetPayload() *models.V1PackReadme {
	return o.Payload
}

func (o *V1PacksUIDReadmeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PackReadme)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

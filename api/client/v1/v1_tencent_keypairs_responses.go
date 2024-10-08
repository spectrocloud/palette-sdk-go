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

// V1TencentKeypairsReader is a Reader for the V1TencentKeypairs structure.
type V1TencentKeypairsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TencentKeypairsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1TencentKeypairsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TencentKeypairsOK creates a V1TencentKeypairsOK with default headers values
func NewV1TencentKeypairsOK() *V1TencentKeypairsOK {
	return &V1TencentKeypairsOK{}
}

/*
V1TencentKeypairsOK handles this case with default header values.

(empty)
*/
type V1TencentKeypairsOK struct {
	Payload *models.V1TencentKeypairs
}

func (o *V1TencentKeypairsOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/tencent/regions/{region}/keypairs][%d] v1TencentKeypairsOK  %+v", 200, o.Payload)
}

func (o *V1TencentKeypairsOK) GetPayload() *models.V1TencentKeypairs {
	return o.Payload
}

func (o *V1TencentKeypairsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1TencentKeypairs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

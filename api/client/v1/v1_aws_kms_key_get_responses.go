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

// V1AwsKmsKeyGetReader is a Reader for the V1AwsKmsKeyGet structure.
type V1AwsKmsKeyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsKmsKeyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AwsKmsKeyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsKmsKeyGetOK creates a V1AwsKmsKeyGetOK with default headers values
func NewV1AwsKmsKeyGetOK() *V1AwsKmsKeyGetOK {
	return &V1AwsKmsKeyGetOK{}
}

/*V1AwsKmsKeyGetOK handles this case with default header values.

(empty)
*/
type V1AwsKmsKeyGetOK struct {
	Payload *models.V1AwsKmsKeyEntity
}

func (o *V1AwsKmsKeyGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/aws/regions/{region}/kms/{keyId}][%d] v1AwsKmsKeyGetOK  %+v", 200, o.Payload)
}

func (o *V1AwsKmsKeyGetOK) GetPayload() *models.V1AwsKmsKeyEntity {
	return o.Payload
}

func (o *V1AwsKmsKeyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AwsKmsKeyEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

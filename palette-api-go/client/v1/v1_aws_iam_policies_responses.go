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

// V1AwsIamPoliciesReader is a Reader for the V1AwsIamPolicies structure.
type V1AwsIamPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsIamPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AwsIamPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsIamPoliciesOK creates a V1AwsIamPoliciesOK with default headers values
func NewV1AwsIamPoliciesOK() *V1AwsIamPoliciesOK {
	return &V1AwsIamPoliciesOK{}
}

/*V1AwsIamPoliciesOK handles this case with default header values.

(empty)
*/
type V1AwsIamPoliciesOK struct {
	Payload *models.V1AwsPolicies
}

func (o *V1AwsIamPoliciesOK) Error() string {
	return fmt.Sprintf("[POST /v1/clouds/aws/policies][%d] v1AwsIamPoliciesOK  %+v", 200, o.Payload)
}

func (o *V1AwsIamPoliciesOK) GetPayload() *models.V1AwsPolicies {
	return o.Payload
}

func (o *V1AwsIamPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AwsPolicies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1AwsSecurityGroupsReader is a Reader for the V1AwsSecurityGroups structure.
type V1AwsSecurityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsSecurityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1AwsSecurityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsSecurityGroupsOK creates a V1AwsSecurityGroupsOK with default headers values
func NewV1AwsSecurityGroupsOK() *V1AwsSecurityGroupsOK {
	return &V1AwsSecurityGroupsOK{}
}

/*
V1AwsSecurityGroupsOK handles this case with default header values.

(empty)
*/
type V1AwsSecurityGroupsOK struct {
	Payload *models.V1AwsSecurityGroups
}

func (o *V1AwsSecurityGroupsOK) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/aws/securitygroups][%d] v1AwsSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *V1AwsSecurityGroupsOK) GetPayload() *models.V1AwsSecurityGroups {
	return o.Payload
}

func (o *V1AwsSecurityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1AwsSecurityGroups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

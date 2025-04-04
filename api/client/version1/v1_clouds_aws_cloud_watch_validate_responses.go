// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CloudsAwsCloudWatchValidateReader is a Reader for the V1CloudsAwsCloudWatchValidate structure.
type V1CloudsAwsCloudWatchValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CloudsAwsCloudWatchValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CloudsAwsCloudWatchValidateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CloudsAwsCloudWatchValidateNoContent creates a V1CloudsAwsCloudWatchValidateNoContent with default headers values
func NewV1CloudsAwsCloudWatchValidateNoContent() *V1CloudsAwsCloudWatchValidateNoContent {
	return &V1CloudsAwsCloudWatchValidateNoContent{}
}

/*
V1CloudsAwsCloudWatchValidateNoContent handles this case with default header values.

Ok response without content
*/
type V1CloudsAwsCloudWatchValidateNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1CloudsAwsCloudWatchValidateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/clouds/aws/cloudwatch/validate][%d] v1CloudsAwsCloudWatchValidateNoContent ", 204)
}

func (o *V1CloudsAwsCloudWatchValidateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}

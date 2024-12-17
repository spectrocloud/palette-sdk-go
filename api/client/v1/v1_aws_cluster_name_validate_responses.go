// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1AwsClusterNameValidateReader is a Reader for the V1AwsClusterNameValidate structure.
type V1AwsClusterNameValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AwsClusterNameValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1AwsClusterNameValidateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AwsClusterNameValidateNoContent creates a V1AwsClusterNameValidateNoContent with default headers values
func NewV1AwsClusterNameValidateNoContent() *V1AwsClusterNameValidateNoContent {
	return &V1AwsClusterNameValidateNoContent{}
}

/*V1AwsClusterNameValidateNoContent handles this case with default header values.

Ok response without content
*/
type V1AwsClusterNameValidateNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1AwsClusterNameValidateNoContent) Error() string {
	return fmt.Sprintf("[GET /v1/clouds/aws/regions/{region}/eksClusters/name/validate][%d] v1AwsClusterNameValidateNoContent ", 204)
}

func (o *V1AwsClusterNameValidateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}

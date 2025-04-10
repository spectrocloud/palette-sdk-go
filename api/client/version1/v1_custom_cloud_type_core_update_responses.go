// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CustomCloudTypeCoreUpdateReader is a Reader for the V1CustomCloudTypeCoreUpdate structure.
type V1CustomCloudTypeCoreUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CustomCloudTypeCoreUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CustomCloudTypeCoreUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CustomCloudTypeCoreUpdateNoContent creates a V1CustomCloudTypeCoreUpdateNoContent with default headers values
func NewV1CustomCloudTypeCoreUpdateNoContent() *V1CustomCloudTypeCoreUpdateNoContent {
	return &V1CustomCloudTypeCoreUpdateNoContent{}
}

/*
V1CustomCloudTypeCoreUpdateNoContent handles this case with default header values.

Ok response without content
*/
type V1CustomCloudTypeCoreUpdateNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1CustomCloudTypeCoreUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/clouds/cloudTypes/{cloudType}/content/core][%d] v1CustomCloudTypeCoreUpdateNoContent ", 204)
}

func (o *V1CustomCloudTypeCoreUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}

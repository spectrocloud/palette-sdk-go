// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1CustomCloudTypeControlPlanePoolTemplateUpdateReader is a Reader for the V1CustomCloudTypeControlPlanePoolTemplateUpdate structure.
type V1CustomCloudTypeControlPlanePoolTemplateUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent creates a V1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent with default headers values
func NewV1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent() *V1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent {
	return &V1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent{}
}

/*
V1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent handles this case with default header values.

Ok response without content
*/
type V1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/clouds/cloudTypes/{cloudType}/content/templates/controlPlanePoolTemplate][%d] v1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent ", 204)
}

func (o *V1CustomCloudTypeControlPlanePoolTemplateUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}

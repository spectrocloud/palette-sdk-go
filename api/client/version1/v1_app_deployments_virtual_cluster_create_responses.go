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

// V1AppDeploymentsVirtualClusterCreateReader is a Reader for the V1AppDeploymentsVirtualClusterCreate structure.
type V1AppDeploymentsVirtualClusterCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1AppDeploymentsVirtualClusterCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewV1AppDeploymentsVirtualClusterCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1AppDeploymentsVirtualClusterCreateCreated creates a V1AppDeploymentsVirtualClusterCreateCreated with default headers values
func NewV1AppDeploymentsVirtualClusterCreateCreated() *V1AppDeploymentsVirtualClusterCreateCreated {
	return &V1AppDeploymentsVirtualClusterCreateCreated{}
}

/*
V1AppDeploymentsVirtualClusterCreateCreated handles this case with default header values.

Created successfully
*/
type V1AppDeploymentsVirtualClusterCreateCreated struct {
	/*Audit uid for the request
	 */
	AuditUID string

	Payload *models.V1UID
}

func (o *V1AppDeploymentsVirtualClusterCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1/appDeployments][%d] v1AppDeploymentsVirtualClusterCreateCreated  %+v", 201, o.Payload)
}

func (o *V1AppDeploymentsVirtualClusterCreateCreated) GetPayload() *models.V1UID {
	return o.Payload
}

func (o *V1AppDeploymentsVirtualClusterCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	o.Payload = new(models.V1UID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

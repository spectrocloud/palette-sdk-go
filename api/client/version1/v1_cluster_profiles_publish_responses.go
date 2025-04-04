// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1ClusterProfilesPublishReader is a Reader for the V1ClusterProfilesPublish structure.
type V1ClusterProfilesPublishReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterProfilesPublishReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1ClusterProfilesPublishNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterProfilesPublishNoContent creates a V1ClusterProfilesPublishNoContent with default headers values
func NewV1ClusterProfilesPublishNoContent() *V1ClusterProfilesPublishNoContent {
	return &V1ClusterProfilesPublishNoContent{}
}

/*
V1ClusterProfilesPublishNoContent handles this case with default header values.

Cluster profile published successfully
*/
type V1ClusterProfilesPublishNoContent struct {
}

func (o *V1ClusterProfilesPublishNoContent) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusterprofiles/{uid}/publish][%d] v1ClusterProfilesPublishNoContent ", 204)
}

func (o *V1ClusterProfilesPublishNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

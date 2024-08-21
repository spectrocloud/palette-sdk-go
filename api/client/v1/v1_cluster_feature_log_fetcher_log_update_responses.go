// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1ClusterFeatureLogFetcherLogUpdateReader is a Reader for the V1ClusterFeatureLogFetcherLogUpdate structure.
type V1ClusterFeatureLogFetcherLogUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterFeatureLogFetcherLogUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1ClusterFeatureLogFetcherLogUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterFeatureLogFetcherLogUpdateNoContent creates a V1ClusterFeatureLogFetcherLogUpdateNoContent with default headers values
func NewV1ClusterFeatureLogFetcherLogUpdateNoContent() *V1ClusterFeatureLogFetcherLogUpdateNoContent {
	return &V1ClusterFeatureLogFetcherLogUpdateNoContent{}
}

/*
V1ClusterFeatureLogFetcherLogUpdateNoContent handles this case with default header values.

Ok response without content
*/
type V1ClusterFeatureLogFetcherLogUpdateNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1ClusterFeatureLogFetcherLogUpdateNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/features/logFetcher/{uid}/log][%d] v1ClusterFeatureLogFetcherLogUpdateNoContent ", 204)
}

func (o *V1ClusterFeatureLogFetcherLogUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}
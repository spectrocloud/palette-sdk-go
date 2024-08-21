// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDWorkloadsKindSyncReader is a Reader for the V1SpectroClustersUIDWorkloadsKindSync structure.
type V1SpectroClustersUIDWorkloadsKindSyncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDWorkloadsKindSyncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewV1SpectroClustersUIDWorkloadsKindSyncAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDWorkloadsKindSyncAccepted creates a V1SpectroClustersUIDWorkloadsKindSyncAccepted with default headers values
func NewV1SpectroClustersUIDWorkloadsKindSyncAccepted() *V1SpectroClustersUIDWorkloadsKindSyncAccepted {
	return &V1SpectroClustersUIDWorkloadsKindSyncAccepted{}
}

/*
V1SpectroClustersUIDWorkloadsKindSyncAccepted handles this case with default header values.

Ok response without content
*/
type V1SpectroClustersUIDWorkloadsKindSyncAccepted struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1SpectroClustersUIDWorkloadsKindSyncAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/{uid}/workloads/{workloadKind}/sync][%d] v1SpectroClustersUidWorkloadsKindSyncAccepted ", 202)
}

func (o *V1SpectroClustersUIDWorkloadsKindSyncAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}
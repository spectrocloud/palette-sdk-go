// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDStatusSpcApplyReader is a Reader for the V1SpectroClustersUIDStatusSpcApply structure.
type V1SpectroClustersUIDStatusSpcApplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDStatusSpcApplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewV1SpectroClustersUIDStatusSpcApplyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDStatusSpcApplyAccepted creates a V1SpectroClustersUIDStatusSpcApplyAccepted with default headers values
func NewV1SpectroClustersUIDStatusSpcApplyAccepted() *V1SpectroClustersUIDStatusSpcApplyAccepted {
	return &V1SpectroClustersUIDStatusSpcApplyAccepted{}
}

/*
V1SpectroClustersUIDStatusSpcApplyAccepted handles this case with default header values.

Ok response without content
*/
type V1SpectroClustersUIDStatusSpcApplyAccepted struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1SpectroClustersUIDStatusSpcApplyAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/{uid}/status/spcApply][%d] v1SpectroClustersUidStatusSpcApplyAccepted ", 202)
}

func (o *V1SpectroClustersUIDStatusSpcApplyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}

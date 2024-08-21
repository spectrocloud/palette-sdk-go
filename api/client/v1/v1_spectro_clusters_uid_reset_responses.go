// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1SpectroClustersUIDResetReader is a Reader for the V1SpectroClustersUIDReset structure.
type V1SpectroClustersUIDResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SpectroClustersUIDResetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDResetNoContent creates a V1SpectroClustersUIDResetNoContent with default headers values
func NewV1SpectroClustersUIDResetNoContent() *V1SpectroClustersUIDResetNoContent {
	return &V1SpectroClustersUIDResetNoContent{}
}

/*
V1SpectroClustersUIDResetNoContent handles this case with default header values.

Ok response without content
*/
type V1SpectroClustersUIDResetNoContent struct {
	/*Audit uid for the request
	 */
	AuditUID string
}

func (o *V1SpectroClustersUIDResetNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/reset][%d] v1SpectroClustersUidResetNoContent ", 204)
}

func (o *V1SpectroClustersUIDResetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header AuditUid
	o.AuditUID = response.GetHeader("AuditUid")

	return nil
}
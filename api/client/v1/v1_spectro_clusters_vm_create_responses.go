// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

// V1SpectroClustersVMCreateReader is a Reader for the V1SpectroClustersVMCreate structure.
type V1SpectroClustersVMCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersVMCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersVMCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersVMCreateOK creates a V1SpectroClustersVMCreateOK with default headers values
func NewV1SpectroClustersVMCreateOK() *V1SpectroClustersVMCreateOK {
	return &V1SpectroClustersVMCreateOK{}
}

/*
V1SpectroClustersVMCreateOK handles this case with default header values.

(empty)
*/
type V1SpectroClustersVMCreateOK struct {
	Payload *models.V1ClusterVirtualMachine
}

func (o *V1SpectroClustersVMCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/{uid}/vms][%d] v1SpectroClustersVmCreateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersVMCreateOK) GetPayload() *models.V1ClusterVirtualMachine {
	return o.Payload
}

func (o *V1SpectroClustersVMCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterVirtualMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
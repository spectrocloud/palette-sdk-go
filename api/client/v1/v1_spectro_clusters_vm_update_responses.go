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

// V1SpectroClustersVMUpdateReader is a Reader for the V1SpectroClustersVMUpdate structure.
type V1SpectroClustersVMUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersVMUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersVMUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersVMUpdateOK creates a V1SpectroClustersVMUpdateOK with default headers values
func NewV1SpectroClustersVMUpdateOK() *V1SpectroClustersVMUpdateOK {
	return &V1SpectroClustersVMUpdateOK{}
}

/*
V1SpectroClustersVMUpdateOK handles this case with default header values.

(empty)
*/
type V1SpectroClustersVMUpdateOK struct {
	Payload *models.V1ClusterVirtualMachine
}

func (o *V1SpectroClustersVMUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/vms/{vmName}][%d] v1SpectroClustersVmUpdateOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersVMUpdateOK) GetPayload() *models.V1ClusterVirtualMachine {
	return o.Payload
}

func (o *V1SpectroClustersVMUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterVirtualMachine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

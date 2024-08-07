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

// V1VMSnapshotUpdateReader is a Reader for the V1VMSnapshotUpdate structure.
type V1VMSnapshotUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1VMSnapshotUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1VMSnapshotUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1VMSnapshotUpdateOK creates a V1VMSnapshotUpdateOK with default headers values
func NewV1VMSnapshotUpdateOK() *V1VMSnapshotUpdateOK {
	return &V1VMSnapshotUpdateOK{}
}

/*
V1VMSnapshotUpdateOK handles this case with default header values.

(empty)
*/
type V1VMSnapshotUpdateOK struct {
	Payload *models.V1VirtualMachineSnapshot
}

func (o *V1VMSnapshotUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/spectroclusters/{uid}/vms/{vmName}/snapshot/{snapshotName}][%d] v1VmSnapshotUpdateOK  %+v", 200, o.Payload)
}

func (o *V1VMSnapshotUpdateOK) GetPayload() *models.V1VirtualMachineSnapshot {
	return o.Payload
}

func (o *V1VMSnapshotUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1VirtualMachineSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

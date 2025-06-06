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

// V1ClusterVMSnapshotsListReader is a Reader for the V1ClusterVMSnapshotsList structure.
type V1ClusterVMSnapshotsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterVMSnapshotsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterVMSnapshotsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterVMSnapshotsListOK creates a V1ClusterVMSnapshotsListOK with default headers values
func NewV1ClusterVMSnapshotsListOK() *V1ClusterVMSnapshotsListOK {
	return &V1ClusterVMSnapshotsListOK{}
}

/*
V1ClusterVMSnapshotsListOK handles this case with default header values.

OK
*/
type V1ClusterVMSnapshotsListOK struct {
	Payload *models.V1VirtualMachineSnapshotList
}

func (o *V1ClusterVMSnapshotsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/vms/snapshot][%d] v1ClusterVmSnapshotsListOK  %+v", 200, o.Payload)
}

func (o *V1ClusterVMSnapshotsListOK) GetPayload() *models.V1VirtualMachineSnapshotList {
	return o.Payload
}

func (o *V1ClusterVMSnapshotsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1VirtualMachineSnapshotList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

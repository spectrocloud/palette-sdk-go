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

// V1LibvirtClustersHostsListReader is a Reader for the V1LibvirtClustersHostsList structure.
type V1LibvirtClustersHostsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1LibvirtClustersHostsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1LibvirtClustersHostsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1LibvirtClustersHostsListOK creates a V1LibvirtClustersHostsListOK with default headers values
func NewV1LibvirtClustersHostsListOK() *V1LibvirtClustersHostsListOK {
	return &V1LibvirtClustersHostsListOK{}
}

/*
V1LibvirtClustersHostsListOK handles this case with default header values.

List of edge host devices
*/
type V1LibvirtClustersHostsListOK struct {
	Payload *models.V1EdgeHostDevices
}

func (o *V1LibvirtClustersHostsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/libvirt/edgeHosts][%d] v1LibvirtClustersHostsListOK  %+v", 200, o.Payload)
}

func (o *V1LibvirtClustersHostsListOK) GetPayload() *models.V1EdgeHostDevices {
	return o.Payload
}

func (o *V1LibvirtClustersHostsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EdgeHostDevices)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

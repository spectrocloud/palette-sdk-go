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

// V1SpectroClustersUIDStatusReader is a Reader for the V1SpectroClustersUIDStatus structure.
type V1SpectroClustersUIDStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersUIDStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDStatusOK creates a V1SpectroClustersUIDStatusOK with default headers values
func NewV1SpectroClustersUIDStatusOK() *V1SpectroClustersUIDStatusOK {
	return &V1SpectroClustersUIDStatusOK{}
}

/*
V1SpectroClustersUIDStatusOK handles this case with default header values.

OK
*/
type V1SpectroClustersUIDStatusOK struct {
	Payload *models.V1SpectroClusterStatusEntity
}

func (o *V1SpectroClustersUIDStatusOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/status][%d] v1SpectroClustersUidStatusOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersUIDStatusOK) GetPayload() *models.V1SpectroClusterStatusEntity {
	return o.Payload
}

func (o *V1SpectroClustersUIDStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterStatusEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
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

// V1SpectroClustersUIDRepaveGetReader is a Reader for the V1SpectroClustersUIDRepaveGet structure.
type V1SpectroClustersUIDRepaveGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDRepaveGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersUIDRepaveGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDRepaveGetOK creates a V1SpectroClustersUIDRepaveGetOK with default headers values
func NewV1SpectroClustersUIDRepaveGetOK() *V1SpectroClustersUIDRepaveGetOK {
	return &V1SpectroClustersUIDRepaveGetOK{}
}

/*
V1SpectroClustersUIDRepaveGetOK handles this case with default header values.

Returns cluster repave status
*/
type V1SpectroClustersUIDRepaveGetOK struct {
	Payload *models.V1SpectroClusterRepave
}

func (o *V1SpectroClustersUIDRepaveGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/repave/status][%d] v1SpectroClustersUidRepaveGetOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersUIDRepaveGetOK) GetPayload() *models.V1SpectroClusterRepave {
	return o.Payload
}

func (o *V1SpectroClustersUIDRepaveGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterRepave)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

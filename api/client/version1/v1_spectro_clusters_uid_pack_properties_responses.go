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

// V1SpectroClustersUIDPackPropertiesReader is a Reader for the V1SpectroClustersUIDPackProperties structure.
type V1SpectroClustersUIDPackPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDPackPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersUIDPackPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDPackPropertiesOK creates a V1SpectroClustersUIDPackPropertiesOK with default headers values
func NewV1SpectroClustersUIDPackPropertiesOK() *V1SpectroClustersUIDPackPropertiesOK {
	return &V1SpectroClustersUIDPackPropertiesOK{}
}

/*
V1SpectroClustersUIDPackPropertiesOK handles this case with default header values.

Cluster's pack properties response
*/
type V1SpectroClustersUIDPackPropertiesOK struct {
	Payload *models.V1SpectroClusterPackProperties
}

func (o *V1SpectroClustersUIDPackPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/pack/properties][%d] v1SpectroClustersUidPackPropertiesOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersUIDPackPropertiesOK) GetPayload() *models.V1SpectroClusterPackProperties {
	return o.Payload
}

func (o *V1SpectroClustersUIDPackPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterPackProperties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

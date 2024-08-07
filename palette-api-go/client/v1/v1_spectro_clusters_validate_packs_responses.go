// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1SpectroClustersValidatePacksReader is a Reader for the V1SpectroClustersValidatePacks structure.
type V1SpectroClustersValidatePacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersValidatePacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersValidatePacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersValidatePacksOK creates a V1SpectroClustersValidatePacksOK with default headers values
func NewV1SpectroClustersValidatePacksOK() *V1SpectroClustersValidatePacksOK {
	return &V1SpectroClustersValidatePacksOK{}
}

/*V1SpectroClustersValidatePacksOK handles this case with default header values.

Cluster packs validation response
*/
type V1SpectroClustersValidatePacksOK struct {
	Payload *models.V1SpectroClusterValidatorResponse
}

func (o *V1SpectroClustersValidatePacksOK) Error() string {
	return fmt.Sprintf("[POST /v1/spectroclusters/validate/packs][%d] v1SpectroClustersValidatePacksOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersValidatePacksOK) GetPayload() *models.V1SpectroClusterValidatorResponse {
	return o.Payload
}

func (o *V1SpectroClustersValidatePacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterValidatorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

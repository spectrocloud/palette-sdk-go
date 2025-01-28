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

// V1SpectroClustersTagsGetReader is a Reader for the V1SpectroClustersTagsGet structure.
type V1SpectroClustersTagsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersTagsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersTagsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersTagsGetOK creates a V1SpectroClustersTagsGetOK with default headers values
func NewV1SpectroClustersTagsGetOK() *V1SpectroClustersTagsGetOK {
	return &V1SpectroClustersTagsGetOK{}
}

/*V1SpectroClustersTagsGetOK handles this case with default header values.

An array of spectrocluster tags
*/
type V1SpectroClustersTagsGetOK struct {
	Payload *models.V1SpectroClusterTags
}

func (o *V1SpectroClustersTagsGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/tags][%d] v1SpectroClustersTagsGetOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersTagsGetOK) GetPayload() *models.V1SpectroClusterTags {
	return o.Payload
}

func (o *V1SpectroClustersTagsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterTags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

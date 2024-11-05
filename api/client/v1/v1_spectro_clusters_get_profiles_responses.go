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

// V1SpectroClustersGetProfilesReader is a Reader for the V1SpectroClustersGetProfiles structure.
type V1SpectroClustersGetProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersGetProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersGetProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersGetProfilesOK creates a V1SpectroClustersGetProfilesOK with default headers values
func NewV1SpectroClustersGetProfilesOK() *V1SpectroClustersGetProfilesOK {
	return &V1SpectroClustersGetProfilesOK{}
}

/*V1SpectroClustersGetProfilesOK handles this case with default header values.

OK
*/
type V1SpectroClustersGetProfilesOK struct {
	Payload *models.V1SpectroClusterProfileList
}

func (o *V1SpectroClustersGetProfilesOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/profiles][%d] v1SpectroClustersGetProfilesOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersGetProfilesOK) GetPayload() *models.V1SpectroClusterProfileList {
	return o.Payload
}

func (o *V1SpectroClustersGetProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterProfileList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

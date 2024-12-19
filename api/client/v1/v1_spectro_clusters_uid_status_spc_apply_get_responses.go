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

// V1SpectroClustersUIDStatusSpcApplyGetReader is a Reader for the V1SpectroClustersUIDStatusSpcApplyGet structure.
type V1SpectroClustersUIDStatusSpcApplyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDStatusSpcApplyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersUIDStatusSpcApplyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDStatusSpcApplyGetOK creates a V1SpectroClustersUIDStatusSpcApplyGetOK with default headers values
func NewV1SpectroClustersUIDStatusSpcApplyGetOK() *V1SpectroClustersUIDStatusSpcApplyGetOK {
	return &V1SpectroClustersUIDStatusSpcApplyGetOK{}
}

/*
V1SpectroClustersUIDStatusSpcApplyGetOK handles this case with default header values.

(empty)
*/
type V1SpectroClustersUIDStatusSpcApplyGetOK struct {
	Payload *models.V1SpcApply
}

func (o *V1SpectroClustersUIDStatusSpcApplyGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/status/spcApply][%d] v1SpectroClustersUidStatusSpcApplyGetOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersUIDStatusSpcApplyGetOK) GetPayload() *models.V1SpcApply {
	return o.Payload
}

func (o *V1SpectroClustersUIDStatusSpcApplyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpcApply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

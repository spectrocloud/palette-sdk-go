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

// V1SpectroClustersUIDKubeCtlRedirectReader is a Reader for the V1SpectroClustersUIDKubeCtlRedirect structure.
type V1SpectroClustersUIDKubeCtlRedirectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDKubeCtlRedirectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersUIDKubeCtlRedirectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDKubeCtlRedirectOK creates a V1SpectroClustersUIDKubeCtlRedirectOK with default headers values
func NewV1SpectroClustersUIDKubeCtlRedirectOK() *V1SpectroClustersUIDKubeCtlRedirectOK {
	return &V1SpectroClustersUIDKubeCtlRedirectOK{}
}

/*
V1SpectroClustersUIDKubeCtlRedirectOK handles this case with default header values.

(empty)
*/
type V1SpectroClustersUIDKubeCtlRedirectOK struct {
	Payload *models.V1SpectroClusterKubeCtlRedirect
}

func (o *V1SpectroClustersUIDKubeCtlRedirectOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/kubectl/redirect][%d] v1SpectroClustersUidKubeCtlRedirectOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersUIDKubeCtlRedirectOK) GetPayload() *models.V1SpectroClusterKubeCtlRedirect {
	return o.Payload
}

func (o *V1SpectroClustersUIDKubeCtlRedirectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterKubeCtlRedirect)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

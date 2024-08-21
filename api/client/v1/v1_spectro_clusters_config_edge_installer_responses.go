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

// V1SpectroClustersConfigEdgeInstallerReader is a Reader for the V1SpectroClustersConfigEdgeInstaller structure.
type V1SpectroClustersConfigEdgeInstallerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersConfigEdgeInstallerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersConfigEdgeInstallerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersConfigEdgeInstallerOK creates a V1SpectroClustersConfigEdgeInstallerOK with default headers values
func NewV1SpectroClustersConfigEdgeInstallerOK() *V1SpectroClustersConfigEdgeInstallerOK {
	return &V1SpectroClustersConfigEdgeInstallerOK{}
}

/*
V1SpectroClustersConfigEdgeInstallerOK handles this case with default header values.

(empty)
*/
type V1SpectroClustersConfigEdgeInstallerOK struct {
	Payload *models.V1ClusterEdgeInstallerConfig
}

func (o *V1SpectroClustersConfigEdgeInstallerOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/config/edgeInstaller][%d] v1SpectroClustersConfigEdgeInstallerOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersConfigEdgeInstallerOK) GetPayload() *models.V1ClusterEdgeInstallerConfig {
	return o.Payload
}

func (o *V1SpectroClustersConfigEdgeInstallerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterEdgeInstallerConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
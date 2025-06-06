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

// V1SpectroClustersMetadataSearchReader is a Reader for the V1SpectroClustersMetadataSearch structure.
type V1SpectroClustersMetadataSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersMetadataSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersMetadataSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersMetadataSearchOK creates a V1SpectroClustersMetadataSearchOK with default headers values
func NewV1SpectroClustersMetadataSearchOK() *V1SpectroClustersMetadataSearchOK {
	return &V1SpectroClustersMetadataSearchOK{}
}

/*
V1SpectroClustersMetadataSearchOK handles this case with default header values.

An array of cluster summary meta items
*/
type V1SpectroClustersMetadataSearchOK struct {
	Payload *models.V1SpectroClustersMetadataSearch
}

func (o *V1SpectroClustersMetadataSearchOK) Error() string {
	return fmt.Sprintf("[POST /v1/dashboard/spectroclusters/metadata/search][%d] v1SpectroClustersMetadataSearchOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersMetadataSearchOK) GetPayload() *models.V1SpectroClustersMetadataSearch {
	return o.Payload
}

func (o *V1SpectroClustersMetadataSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClustersMetadataSearch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

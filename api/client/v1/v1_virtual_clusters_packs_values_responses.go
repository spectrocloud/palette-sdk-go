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

// V1VirtualClustersPacksValuesReader is a Reader for the V1VirtualClustersPacksValues structure.
type V1VirtualClustersPacksValuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1VirtualClustersPacksValuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1VirtualClustersPacksValuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1VirtualClustersPacksValuesOK creates a V1VirtualClustersPacksValuesOK with default headers values
func NewV1VirtualClustersPacksValuesOK() *V1VirtualClustersPacksValuesOK {
	return &V1VirtualClustersPacksValuesOK{}
}

/*
V1VirtualClustersPacksValuesOK handles this case with default header values.

Successful response
*/
type V1VirtualClustersPacksValuesOK struct {
	Payload *models.V1ClusterVirtualPacksValues
}

func (o *V1VirtualClustersPacksValuesOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/virtual/packs/values][%d] v1VirtualClustersPacksValuesOK  %+v", 200, o.Payload)
}

func (o *V1VirtualClustersPacksValuesOK) GetPayload() *models.V1ClusterVirtualPacksValues {
	return o.Payload
}

func (o *V1VirtualClustersPacksValuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterVirtualPacksValues)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

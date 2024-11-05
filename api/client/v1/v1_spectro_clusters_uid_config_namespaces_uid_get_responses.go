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

// V1SpectroClustersUIDConfigNamespacesUIDGetReader is a Reader for the V1SpectroClustersUIDConfigNamespacesUIDGet structure.
type V1SpectroClustersUIDConfigNamespacesUIDGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDConfigNamespacesUIDGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersUIDConfigNamespacesUIDGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDConfigNamespacesUIDGetOK creates a V1SpectroClustersUIDConfigNamespacesUIDGetOK with default headers values
func NewV1SpectroClustersUIDConfigNamespacesUIDGetOK() *V1SpectroClustersUIDConfigNamespacesUIDGetOK {
	return &V1SpectroClustersUIDConfigNamespacesUIDGetOK{}
}

/*
V1SpectroClustersUIDConfigNamespacesUIDGetOK handles this case with default header values.

Cluster's namespace response
*/
type V1SpectroClustersUIDConfigNamespacesUIDGetOK struct {
	Payload *models.V1ClusterNamespaceResource
}

func (o *V1SpectroClustersUIDConfigNamespacesUIDGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/config/namespaces/{namespaceUid}][%d] v1SpectroClustersUidConfigNamespacesUidGetOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersUIDConfigNamespacesUIDGetOK) GetPayload() *models.V1ClusterNamespaceResource {
	return o.Payload
}

func (o *V1SpectroClustersUIDConfigNamespacesUIDGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterNamespaceResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

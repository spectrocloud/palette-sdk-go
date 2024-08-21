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

// V1SpectroClustersUIDOIDCDashboardURLReader is a Reader for the V1SpectroClustersUIDOIDCDashboardURL structure.
type V1SpectroClustersUIDOIDCDashboardURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDOIDCDashboardURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersUIDOIDCDashboardURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDOIDCDashboardURLOK creates a V1SpectroClustersUIDOIDCDashboardURLOK with default headers values
func NewV1SpectroClustersUIDOIDCDashboardURLOK() *V1SpectroClustersUIDOIDCDashboardURLOK {
	return &V1SpectroClustersUIDOIDCDashboardURLOK{}
}

/*
V1SpectroClustersUIDOIDCDashboardURLOK handles this case with default header values.

OK
*/
type V1SpectroClustersUIDOIDCDashboardURLOK struct {
	Payload *models.V1SectroClusterK8sDashboardURL
}

func (o *V1SpectroClustersUIDOIDCDashboardURLOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/oidc/dashboard/url][%d] v1SpectroClustersUidOIdCDashboardUrlOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersUIDOIDCDashboardURLOK) GetPayload() *models.V1SectroClusterK8sDashboardURL {
	return o.Payload
}

func (o *V1SpectroClustersUIDOIDCDashboardURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SectroClusterK8sDashboardURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
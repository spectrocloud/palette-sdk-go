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

// V1TenantUIDAssetsCertsListReader is a Reader for the V1TenantUIDAssetsCertsList structure.
type V1TenantUIDAssetsCertsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantUIDAssetsCertsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1TenantUIDAssetsCertsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantUIDAssetsCertsListOK creates a V1TenantUIDAssetsCertsListOK with default headers values
func NewV1TenantUIDAssetsCertsListOK() *V1TenantUIDAssetsCertsListOK {
	return &V1TenantUIDAssetsCertsListOK{}
}

/*
V1TenantUIDAssetsCertsListOK handles this case with default header values.

OK
*/
type V1TenantUIDAssetsCertsListOK struct {
	Payload *models.V1TenantAssetCerts
}

func (o *V1TenantUIDAssetsCertsListOK) Error() string {
	return fmt.Sprintf("[GET /v1/tenants/{tenantUid}/assets/certs][%d] v1TenantUIdAssetsCertsListOK  %+v", 200, o.Payload)
}

func (o *V1TenantUIDAssetsCertsListOK) GetPayload() *models.V1TenantAssetCerts {
	return o.Payload
}

func (o *V1TenantUIDAssetsCertsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1TenantAssetCerts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

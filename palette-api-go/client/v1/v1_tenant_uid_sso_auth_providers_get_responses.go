// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1TenantUIDSsoAuthProvidersGetReader is a Reader for the V1TenantUIDSsoAuthProvidersGet structure.
type V1TenantUIDSsoAuthProvidersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1TenantUIDSsoAuthProvidersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1TenantUIDSsoAuthProvidersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1TenantUIDSsoAuthProvidersGetOK creates a V1TenantUIDSsoAuthProvidersGetOK with default headers values
func NewV1TenantUIDSsoAuthProvidersGetOK() *V1TenantUIDSsoAuthProvidersGetOK {
	return &V1TenantUIDSsoAuthProvidersGetOK{}
}

/*V1TenantUIDSsoAuthProvidersGetOK handles this case with default header values.

OK
*/
type V1TenantUIDSsoAuthProvidersGetOK struct {
	Payload *models.V1TenantSsoAuthProvidersEntity
}

func (o *V1TenantUIDSsoAuthProvidersGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/tenants/{tenantUid}/sso/auth/providers][%d] v1TenantUidSsoAuthProvidersGetOK  %+v", 200, o.Payload)
}

func (o *V1TenantUIDSsoAuthProvidersGetOK) GetPayload() *models.V1TenantSsoAuthProvidersEntity {
	return o.Payload
}

func (o *V1TenantUIDSsoAuthProvidersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1TenantSsoAuthProvidersEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

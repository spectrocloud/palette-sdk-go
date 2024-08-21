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

// V1MaasAccountsUIDDomainsReader is a Reader for the V1MaasAccountsUIDDomains structure.
type V1MaasAccountsUIDDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1MaasAccountsUIDDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1MaasAccountsUIDDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1MaasAccountsUIDDomainsOK creates a V1MaasAccountsUIDDomainsOK with default headers values
func NewV1MaasAccountsUIDDomainsOK() *V1MaasAccountsUIDDomainsOK {
	return &V1MaasAccountsUIDDomainsOK{}
}

/*
V1MaasAccountsUIDDomainsOK handles this case with default header values.

(empty)
*/
type V1MaasAccountsUIDDomainsOK struct {
	Payload *models.V1MaasDomains
}

func (o *V1MaasAccountsUIDDomainsOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/maas/{uid}/properties/domains][%d] v1MaasAccountsUidDomainsOK  %+v", 200, o.Payload)
}

func (o *V1MaasAccountsUIDDomainsOK) GetPayload() *models.V1MaasDomains {
	return o.Payload
}

func (o *V1MaasAccountsUIDDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MaasDomains)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
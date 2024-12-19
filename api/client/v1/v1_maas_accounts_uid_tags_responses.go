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

// V1MaasAccountsUIDTagsReader is a Reader for the V1MaasAccountsUIDTags structure.
type V1MaasAccountsUIDTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1MaasAccountsUIDTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1MaasAccountsUIDTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1MaasAccountsUIDTagsOK creates a V1MaasAccountsUIDTagsOK with default headers values
func NewV1MaasAccountsUIDTagsOK() *V1MaasAccountsUIDTagsOK {
	return &V1MaasAccountsUIDTagsOK{}
}

/*V1MaasAccountsUIDTagsOK handles this case with default header values.

(empty)
*/
type V1MaasAccountsUIDTagsOK struct {
	Payload *models.V1MaasTags
}

func (o *V1MaasAccountsUIDTagsOK) Error() string {
	return fmt.Sprintf("[GET /v1/cloudaccounts/maas/{uid}/properties/tags][%d] v1MaasAccountsUidTagsOK  %+v", 200, o.Payload)
}

func (o *V1MaasAccountsUIDTagsOK) GetPayload() *models.V1MaasTags {
	return o.Payload
}

func (o *V1MaasAccountsUIDTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MaasTags)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

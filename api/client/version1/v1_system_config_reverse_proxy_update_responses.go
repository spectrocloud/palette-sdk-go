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

// V1SystemConfigReverseProxyUpdateReader is a Reader for the V1SystemConfigReverseProxyUpdate structure.
type V1SystemConfigReverseProxyUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SystemConfigReverseProxyUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1SystemConfigReverseProxyUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SystemConfigReverseProxyUpdateNoContent creates a V1SystemConfigReverseProxyUpdateNoContent with default headers values
func NewV1SystemConfigReverseProxyUpdateNoContent() *V1SystemConfigReverseProxyUpdateNoContent {
	return &V1SystemConfigReverseProxyUpdateNoContent{}
}

/*
V1SystemConfigReverseProxyUpdateNoContent handles this case with default header values.

(empty)
*/
type V1SystemConfigReverseProxyUpdateNoContent struct {
	Payload models.V1Updated
}

func (o *V1SystemConfigReverseProxyUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/system/config/reverseproxy][%d] v1SystemConfigReverseProxyUpdateNoContent  %+v", 204, o.Payload)
}

func (o *V1SystemConfigReverseProxyUpdateNoContent) GetPayload() models.V1Updated {
	return o.Payload
}

func (o *V1SystemConfigReverseProxyUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1SystemConfigReverseProxyGetReader is a Reader for the V1SystemConfigReverseProxyGet structure.
type V1SystemConfigReverseProxyGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1SystemConfigReverseProxyGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SystemConfigReverseProxyGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SystemConfigReverseProxyGetOK creates a V1SystemConfigReverseProxyGetOK with default headers values
func NewV1SystemConfigReverseProxyGetOK() *V1SystemConfigReverseProxyGetOK {
	return &V1SystemConfigReverseProxyGetOK{}
}

/*
V1SystemConfigReverseProxyGetOK handles this case with default header values.

(empty)
*/
type V1SystemConfigReverseProxyGetOK struct {
	Payload *models.V1SystemReverseProxy
}

func (o *V1SystemConfigReverseProxyGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/system/config/reverseproxy][%d] v1SystemConfigReverseProxyGetOK  %+v", 200, o.Payload)
}

func (o *V1SystemConfigReverseProxyGetOK) GetPayload() *models.V1SystemReverseProxy {
	return o.Payload
}

func (o *V1SystemConfigReverseProxyGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SystemReverseProxy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
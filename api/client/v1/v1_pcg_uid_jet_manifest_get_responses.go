// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1PcgUIDJetManifestGetReader is a Reader for the V1PcgUIDJetManifestGet structure.
type V1PcgUIDJetManifestGetReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *V1PcgUIDJetManifestGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1PcgUIDJetManifestGetOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1PcgUIDJetManifestGetOK creates a V1PcgUIDJetManifestGetOK with default headers values
func NewV1PcgUIDJetManifestGetOK(writer io.Writer) *V1PcgUIDJetManifestGetOK {
	return &V1PcgUIDJetManifestGetOK{
		Payload: writer,
	}
}

/*V1PcgUIDJetManifestGetOK handles this case with default header values.

download file
*/
type V1PcgUIDJetManifestGetOK struct {
	ContentDisposition string

	Payload io.Writer
}

func (o *V1PcgUIDJetManifestGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/pcg/{uid}/services/jet/manifest][%d] v1PcgUidJetManifestGetOK  %+v", 200, o.Payload)
}

func (o *V1PcgUIDJetManifestGetOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *V1PcgUIDJetManifestGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Disposition
	o.ContentDisposition = response.GetHeader("Content-Disposition")

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// V1PcgUIDAllyManifestGetReader is a Reader for the V1PcgUIDAllyManifestGet structure.
type V1PcgUIDAllyManifestGetReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *V1PcgUIDAllyManifestGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1PcgUIDAllyManifestGetOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1PcgUIDAllyManifestGetOK creates a V1PcgUIDAllyManifestGetOK with default headers values
func NewV1PcgUIDAllyManifestGetOK(writer io.Writer) *V1PcgUIDAllyManifestGetOK {
	return &V1PcgUIDAllyManifestGetOK{
		Payload: writer,
	}
}

/*V1PcgUIDAllyManifestGetOK handles this case with default header values.

download file
*/
type V1PcgUIDAllyManifestGetOK struct {
	ContentDisposition string

	Payload io.Writer
}

func (o *V1PcgUIDAllyManifestGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/pcg/{uid}/services/ally/manifest][%d] v1PcgUidAllyManifestGetOK  %+v", 200, o.Payload)
}

func (o *V1PcgUIDAllyManifestGetOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *V1PcgUIDAllyManifestGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Disposition
	o.ContentDisposition = response.GetHeader("Content-Disposition")

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

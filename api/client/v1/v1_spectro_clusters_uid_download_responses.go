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

// V1SpectroClustersUIDDownloadReader is a Reader for the V1SpectroClustersUIDDownload structure.
type V1SpectroClustersUIDDownloadReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *V1SpectroClustersUIDDownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1SpectroClustersUIDDownloadOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1SpectroClustersUIDDownloadOK creates a V1SpectroClustersUIDDownloadOK with default headers values
func NewV1SpectroClustersUIDDownloadOK(writer io.Writer) *V1SpectroClustersUIDDownloadOK {
	return &V1SpectroClustersUIDDownloadOK{
		Payload: writer,
	}
}

/*V1SpectroClustersUIDDownloadOK handles this case with default header values.

download cluster archive file
*/
type V1SpectroClustersUIDDownloadOK struct {
	ContentDisposition string

	Payload io.Writer
}

func (o *V1SpectroClustersUIDDownloadOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/download][%d] v1SpectroClustersUidDownloadOK  %+v", 200, o.Payload)
}

func (o *V1SpectroClustersUIDDownloadOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *V1SpectroClustersUIDDownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Disposition
	o.ContentDisposition = response.GetHeader("Content-Disposition")

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

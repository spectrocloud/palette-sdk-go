// Code generated by go-swagger; DO NOT EDIT.

package version1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1ClusterProfilesUIDSpcDownloadReader is a Reader for the V1ClusterProfilesUIDSpcDownload structure.
type V1ClusterProfilesUIDSpcDownloadReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterProfilesUIDSpcDownloadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterProfilesUIDSpcDownloadOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterProfilesUIDSpcDownloadOK creates a V1ClusterProfilesUIDSpcDownloadOK with default headers values
func NewV1ClusterProfilesUIDSpcDownloadOK(writer io.Writer) *V1ClusterProfilesUIDSpcDownloadOK {
	return &V1ClusterProfilesUIDSpcDownloadOK{
		Payload: writer,
	}
}

/*
V1ClusterProfilesUIDSpcDownloadOK handles this case with default header values.

Download cluster profile archive file
*/
type V1ClusterProfilesUIDSpcDownloadOK struct {
	ContentDisposition string

	Payload io.Writer
}

func (o *V1ClusterProfilesUIDSpcDownloadOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusterprofiles/{uid}/spc/download][%d] v1ClusterProfilesUidSpcDownloadOK  %+v", 200, o.Payload)
}

func (o *V1ClusterProfilesUIDSpcDownloadOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *V1ClusterProfilesUIDSpcDownloadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Disposition
	o.ContentDisposition = response.GetHeader("Content-Disposition")

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

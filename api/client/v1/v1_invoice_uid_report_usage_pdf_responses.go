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

// V1InvoiceUIDReportUsagePdfReader is a Reader for the V1InvoiceUIDReportUsagePdf structure.
type V1InvoiceUIDReportUsagePdfReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *V1InvoiceUIDReportUsagePdfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1InvoiceUIDReportUsagePdfOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1InvoiceUIDReportUsagePdfOK creates a V1InvoiceUIDReportUsagePdfOK with default headers values
func NewV1InvoiceUIDReportUsagePdfOK(writer io.Writer) *V1InvoiceUIDReportUsagePdfOK {
	return &V1InvoiceUIDReportUsagePdfOK{
		Payload: writer,
	}
}

/*
V1InvoiceUIDReportUsagePdfOK handles this case with default header values.

OK
*/
type V1InvoiceUIDReportUsagePdfOK struct {
	ContentDisposition string

	ContentType string

	Payload io.Writer
}

func (o *V1InvoiceUIDReportUsagePdfOK) Error() string {
	return fmt.Sprintf("[GET /v1/tenants/{tenantUid}/invoices/{invoiceUid}/report/usage/pdf][%d] v1InvoiceUidReportUsagePdfOK  %+v", 200, o.Payload)
}

func (o *V1InvoiceUIDReportUsagePdfOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *V1InvoiceUIDReportUsagePdfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Content-Disposition
	o.ContentDisposition = response.GetHeader("Content-Disposition")

	// response header Content-Type
	o.ContentType = response.GetHeader("Content-Type")

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

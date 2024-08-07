package transport

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/go-openapi/runtime"

	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func ReadTooManyRequestError(protocol, path string) (interface{}, error) {

	result := NewTransportError()
	result.HttpCode = http.StatusTooManyRequests
	msg := fmt.Sprintf("Too Many Requests, [%v] %v", protocol, path)
	result.Payload = &models.V1Error{
		Code:    "TooManyRequests",
		Details: nil,
		Message: msg,
		Ref:     "",
	}
	return nil, result
}

func ReadErrorResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewTransportError()
	if err := result.readResponse(response, consumer); err != nil {
		return nil, err
	}
	result.HttpCode = response.Code()
	return nil, result
}

func NewTransportError() *TransportError {
	return &TransportError{}
}

type TransportError struct {
	Payload  *models.V1Error
	HttpCode int
}

func (o *TransportError) Error() string {
	return fmt.Sprintf(" %+v", o.Payload)
}

func (o *TransportError) GetPayload() *models.V1Error {
	return o.Payload
}

func (o *TransportError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer) error {
	o.Payload = new(models.V1Error)
	reader := response.Body()
	rawBody, err := io.ReadAll(response.Body())
	if err == nil {
		reader = io.NopCloser(bytes.NewBuffer(rawBody))
	}
	if err := consumer.Consume(reader, o.Payload); err != nil && err != io.EOF {
		o.Payload = &models.V1Error{
			Code:    http.StatusText(response.Code()),
			Details: string(rawBody),
			Message: response.Message(),
			Ref:     "",
		}
	}
	return nil
}

type TcpError struct {
	Op      string
	Url     string
	Code    string
	Message string
}

func NewTcpError() *TcpError {
	return &TcpError{}
}

func (t *TcpError) readResponse(err error) {
	urlError, ok := err.(*url.Error)
	t.Op = urlError.Op
	t.Url = urlError.URL
	if ok {
		opError, ok := urlError.Err.(*net.OpError)
		if ok {
			sysCallError, ok := opError.Err.(*os.SyscallError)
			if ok {
				t.Code = sysCallError.Syscall
				t.Message = sysCallError.Error()
				return
			}
		}

		//BET-2377 : Return error message for Unhandled error types. Otherwise actual err message gets lost
		t.Message = fmt.Sprintf("An Unknown error occurred. %s", urlError.Err)
	}
}

func (t *TcpError) Error() string {
	return fmt.Sprintf(" %v", t.Message)
}

func TcpErrorResponse(err error) error {
	result := NewTcpError()
	result.readResponse(err)
	return result
}

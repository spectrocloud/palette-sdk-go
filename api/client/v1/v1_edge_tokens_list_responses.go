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

// V1EdgeTokensListReader is a Reader for the V1EdgeTokensList structure.
type V1EdgeTokensListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1EdgeTokensListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1EdgeTokensListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1EdgeTokensListOK creates a V1EdgeTokensListOK with default headers values
func NewV1EdgeTokensListOK() *V1EdgeTokensListOK {
	return &V1EdgeTokensListOK{}
}

/*V1EdgeTokensListOK handles this case with default header values.

An array of edge tokens
*/
type V1EdgeTokensListOK struct {
	Payload *models.V1EdgeTokens
}

func (o *V1EdgeTokensListOK) Error() string {
	return fmt.Sprintf("[GET /v1/edgehosts/tokens][%d] v1EdgeTokensListOK  %+v", 200, o.Payload)
}

func (o *V1EdgeTokensListOK) GetPayload() *models.V1EdgeTokens {
	return o.Payload
}

func (o *V1EdgeTokensListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1EdgeTokens)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/spectrocloud/palette-api-go/models"
)

// V1PacksNameRegistryUIDListReader is a Reader for the V1PacksNameRegistryUIDList structure.
type V1PacksNameRegistryUIDListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1PacksNameRegistryUIDListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1PacksNameRegistryUIDListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1PacksNameRegistryUIDListOK creates a V1PacksNameRegistryUIDListOK with default headers values
func NewV1PacksNameRegistryUIDListOK() *V1PacksNameRegistryUIDListOK {
	return &V1PacksNameRegistryUIDListOK{}
}

/*V1PacksNameRegistryUIDListOK handles this case with default header values.

(empty)
*/
type V1PacksNameRegistryUIDListOK struct {
	Payload *models.V1PackTagEntity
}

func (o *V1PacksNameRegistryUIDListOK) Error() string {
	return fmt.Sprintf("[GET /v1/packs/{packName}/registries/{registryUid}][%d] v1PacksNameRegistryUidListOK  %+v", 200, o.Payload)
}

func (o *V1PacksNameRegistryUIDListOK) GetPayload() *models.V1PackTagEntity {
	return o.Payload
}

func (o *V1PacksNameRegistryUIDListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1PackTagEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

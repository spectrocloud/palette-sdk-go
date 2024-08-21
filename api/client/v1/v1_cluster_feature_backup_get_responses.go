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

// V1ClusterFeatureBackupGetReader is a Reader for the V1ClusterFeatureBackupGet structure.
type V1ClusterFeatureBackupGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterFeatureBackupGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterFeatureBackupGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterFeatureBackupGetOK creates a V1ClusterFeatureBackupGetOK with default headers values
func NewV1ClusterFeatureBackupGetOK() *V1ClusterFeatureBackupGetOK {
	return &V1ClusterFeatureBackupGetOK{}
}

/*
V1ClusterFeatureBackupGetOK handles this case with default header values.

OK
*/
type V1ClusterFeatureBackupGetOK struct {
	Payload *models.V1ClusterBackup
}

func (o *V1ClusterFeatureBackupGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/spectroclusters/{uid}/features/backup][%d] v1ClusterFeatureBackupGetOK  %+v", 200, o.Payload)
}

func (o *V1ClusterFeatureBackupGetOK) GetPayload() *models.V1ClusterBackup {
	return o.Payload
}

func (o *V1ClusterFeatureBackupGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterBackup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
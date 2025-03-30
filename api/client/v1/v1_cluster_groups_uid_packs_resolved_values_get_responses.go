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

// V1ClusterGroupsUIDPacksResolvedValuesGetReader is a Reader for the V1ClusterGroupsUIDPacksResolvedValuesGet structure.
type V1ClusterGroupsUIDPacksResolvedValuesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterGroupsUIDPacksResolvedValuesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV1ClusterGroupsUIDPacksResolvedValuesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterGroupsUIDPacksResolvedValuesGetOK creates a V1ClusterGroupsUIDPacksResolvedValuesGetOK with default headers values
func NewV1ClusterGroupsUIDPacksResolvedValuesGetOK() *V1ClusterGroupsUIDPacksResolvedValuesGetOK {
	return &V1ClusterGroupsUIDPacksResolvedValuesGetOK{}
}

/*
V1ClusterGroupsUIDPacksResolvedValuesGetOK handles this case with default header values.

OK
*/
type V1ClusterGroupsUIDPacksResolvedValuesGetOK struct {
	Payload *models.V1SpectroClusterProfilesResolvedValues
}

func (o *V1ClusterGroupsUIDPacksResolvedValuesGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/clustergroups/{uid}/packs/resolvedValues][%d] v1ClusterGroupsUidPacksResolvedValuesGetOK  %+v", 200, o.Payload)
}

func (o *V1ClusterGroupsUIDPacksResolvedValuesGetOK) GetPayload() *models.V1SpectroClusterProfilesResolvedValues {
	return o.Payload
}

func (o *V1ClusterGroupsUIDPacksResolvedValuesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SpectroClusterProfilesResolvedValues)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

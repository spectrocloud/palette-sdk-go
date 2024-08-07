// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// V1ClusterProfilesUIDPacksNameManifestsUIDUpdateReader is a Reader for the V1ClusterProfilesUIDPacksNameManifestsUIDUpdate structure.
type V1ClusterProfilesUIDPacksNameManifestsUIDUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V1ClusterProfilesUIDPacksNameManifestsUIDUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewV1ClusterProfilesUIDPacksNameManifestsUIDUpdateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV1ClusterProfilesUIDPacksNameManifestsUIDUpdateNoContent creates a V1ClusterProfilesUIDPacksNameManifestsUIDUpdateNoContent with default headers values
func NewV1ClusterProfilesUIDPacksNameManifestsUIDUpdateNoContent() *V1ClusterProfilesUIDPacksNameManifestsUIDUpdateNoContent {
	return &V1ClusterProfilesUIDPacksNameManifestsUIDUpdateNoContent{}
}

/*
V1ClusterProfilesUIDPacksNameManifestsUIDUpdateNoContent handles this case with default header values.

The resource was updated successfully
*/
type V1ClusterProfilesUIDPacksNameManifestsUIDUpdateNoContent struct {
}

func (o *V1ClusterProfilesUIDPacksNameManifestsUIDUpdateNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/clusterprofiles/{uid}/packs/{packName}/manifests/{manifestUid}][%d] v1ClusterProfilesUidPacksNameManifestsUidUpdateNoContent ", 204)
}

func (o *V1ClusterProfilesUIDPacksNameManifestsUIDUpdateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

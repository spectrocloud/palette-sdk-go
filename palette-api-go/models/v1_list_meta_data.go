// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1ListMetaData ListMeta describes metadata for the resource listing
//
// swagger:model v1ListMetaData
type V1ListMetaData struct {

	// Next token for the pagination. Next token is equal to empty string indicates end of result set.
	Continue string `json:"continue"`

	// Total count of the resources which might change during pagination based on the resources addition or deletion
	Count int64 `json:"count"`

	// Number of records feteched
	Limit int64 `json:"limit"`

	// The next offset for the pagination. Starting index for which next request will be placed.
	Offset int64 `json:"offset"`
}

// Validate validates this v1 list meta data
func (m *V1ListMetaData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1ListMetaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1ListMetaData) UnmarshalBinary(b []byte) error {
	var res V1ListMetaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

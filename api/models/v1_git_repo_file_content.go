// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1GitRepoFileContent v1 git repo file content
//
// swagger:model v1GitRepoFileContent
type V1GitRepoFileContent struct {

	// content
	Content string `json:"content,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// repo name
	RepoName string `json:"repoName,omitempty"`

	// sha
	Sha string `json:"sha,omitempty"`
}

// Validate validates this v1 git repo file content
func (m *V1GitRepoFileContent) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 git repo file content based on context it is used
func (m *V1GitRepoFileContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1GitRepoFileContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GitRepoFileContent) UnmarshalBinary(b []byte) error {
	var res V1GitRepoFileContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

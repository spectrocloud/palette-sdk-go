// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// V1VMNUMAGuestMappingPassthrough NUMAGuestMappingPassthrough instructs kubevirt to model numa topology which is compatible with the CPU pinning on the guest. This will result in a subset of the node numa topology being passed through, ensuring that virtual numa nodes and their memory never cross boundaries coming from the node numa mapping.
//
// swagger:model v1VmNUMAGuestMappingPassthrough
type V1VMNUMAGuestMappingPassthrough interface{}

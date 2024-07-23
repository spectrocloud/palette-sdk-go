// Package herr provides error handling utilities for the Spectro Cloud API client.
package herr

import "github.com/spectrocloud/palette-sdk-go/client/apiutil"

// IsNotFound returns a boolean indicating whether an error is a ResourceNotFound error.
func IsNotFound(err error) bool {
	return apiutil.ToV1ErrorObj(err).Code == "ResourceNotFound"
}

// IsEdgeHostDeviceNotRegistered returns a boolean indicating whether an error is an EdgeHostDeviceNotRegistered error.
func IsEdgeHostDeviceNotRegistered(err error) bool {
	return apiutil.ToV1ErrorObj(err).Code == "EdgeHostDeviceNotRegistered"
}

// IsBackupNotConfigured returns a boolean indicating whether an error is a BackupNotConfigured error.
func IsBackupNotConfigured(err error) bool {
	return apiutil.ToV1ErrorObj(err).Code == "BackupNotConfigured"
}

package herr

import "github.com/spectrocloud/palette-sdk-go/client/apiutil"

func IsNotFound(err error) bool {
	return apiutil.ToV1ErrorObj(err).Code == "ResourceNotFound"
}

func IsEdgeHostDeviceNotRegistered(err error) bool {
	return apiutil.ToV1ErrorObj(err).Code == "EdgeHostDeviceNotRegistered"
}

func IsBackupNotConfigured(err error) bool {
	return apiutil.ToV1ErrorObj(err).Code == "BackupNotConfigured"
}

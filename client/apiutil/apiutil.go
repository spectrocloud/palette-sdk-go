package apiutil

import (
	"encoding/base64"
	"fmt"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	"github.com/spectrocloud/palette-api-go/models"
)

func IsBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func Base64DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func ToV1ErrorObj(err interface{}) *models.V1Error {
	if err != nil {
		tErr, ok := err.(*transport.TransportError)
		if ok {
			return tErr.Payload
		}

		goErr, ok := err.(error)
		if ok {
			return &models.V1Error{
				Code:    "UnknownError",
				Message: goErr.Error(),
			}
		}

		return &models.V1Error{
			Code:    "UnknownError",
			Message: fmt.Sprintf("%v", err),
		}
	}
	return &models.V1Error{
		Code: "EmptyError",
	}
}

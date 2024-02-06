package apiutil

import (
	"fmt"

	"github.com/spectrocloud/palette-api-go/apiutil/transport"
	"github.com/spectrocloud/palette-api-go/models"
)

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

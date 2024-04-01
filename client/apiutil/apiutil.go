package apiutil

import (
	"encoding/base64"
	"errors"
	"fmt"
	"hash/fnv"
	"strconv"

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

func StringHash(name string) string {
	return strconv.FormatUint(uint64(hash(name)), 10)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

func Ptr[T any](v T) *T {
	return &v
}

func Handle404(err error) error {
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return nil
	} else if err != nil {
		return err
	}
	return nil
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

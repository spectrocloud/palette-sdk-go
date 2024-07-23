// Package apiutil provides utilities for the Spectro Cloud API client.
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

// IsBase64 returns a boolean indicating whether a string is base64 encoded.
func IsBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

// Base64DecodeString decodes a base64-encoded string.
func Base64DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// StringHash hashes a string using FNV-1a.
func StringHash(name string) string {
	return strconv.FormatUint(uint64(hash(name)), 10)
}

func hash(s string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(s))
	return h.Sum32()
}

// Ptr returns a pointer to any value.
func Ptr[T any](v T) *T {
	return &v
}

// Is404 returns a boolean indicating whether an error is a 404 error.
func Is404(err error) bool {
	var e *transport.TransportError
	if errors.As(err, &e) && e.HttpCode == 404 {
		return true
	}
	return false
}

// ToV1ErrorObj converts an error to a V1Error object.
func ToV1ErrorObj(err any) *models.V1Error {
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

package client

import (
	"github.com/spectrocloud/palette-sdk-go/api/models"
)

func and() *models.V1SearchFilterConjunctionOperator {
	and := models.V1SearchFilterConjunctionOperatorAnd
	return &and
}

func or() *models.V1SearchFilterConjunctionOperator {
	or := models.V1SearchFilterConjunctionOperatorOr
	return &or
}

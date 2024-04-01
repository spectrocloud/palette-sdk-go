package client

import (
	"github.com/spectrocloud/palette-api-go/models"
)

func and() *models.V1SearchFilterConjunctionOperator {
	and := models.V1SearchFilterConjunctionOperatorAnd
	return &and
}

func or() *models.V1SearchFilterConjunctionOperator {
	or := models.V1SearchFilterConjunctionOperatorOr
	return &or
}

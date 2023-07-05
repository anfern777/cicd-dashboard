package entity

import (
	"gorm.io/gorm"
)

type CloudProviderIntegration struct {
	gorm.Model
	ProjectID uint
	Provider  string `validate:"required"`
	Key       string
	Secret    string
}

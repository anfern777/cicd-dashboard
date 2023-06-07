package entity

import "gorm.io/gorm"

type CloudProvider string

const (
	AWS CloudProvider = "AWS"
)

type CloudProviderIntegration struct {
	gorm.Model
	Provider CloudProvider `validate:"required"`
	Key      string
	Secret   string `json:"-"`
}

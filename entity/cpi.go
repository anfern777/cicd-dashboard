package entity

import "gorm.io/gorm"

type CloudProvider string
const (		
	AWS CloudProvider = "AWS"
)

type CloudProviderIntegration struct {
	gorm.Model
	Provider CloudProvider `json:"provider"`
	Key string `json:"key"`
	Secret string `json:"-"`
}
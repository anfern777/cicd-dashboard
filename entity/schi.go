package entity

import "gorm.io/gorm"

type SourceCodeHostProvider string
const (
	GitHub SourceCodeHostProvider = "GitHub"
)

type SourceCodeHostIntegration struct {
	gorm.Model
	Provider SourceCodeHostProvider `json:"provider"`
	BaseUrl string `json:"baseUrl"`
	Secret string `json:"-"`
}
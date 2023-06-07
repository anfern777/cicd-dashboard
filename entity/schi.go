package entity

import "gorm.io/gorm"

type SourceCodeHostProvider string

const (
	GitHub SourceCodeHostProvider = "GitHub"
)

type SourceCodeHostIntegration struct {
	gorm.Model
	Provider SourceCodeHostProvider `validate:"required"`
	BaseUrl  string                 `validate:"required,url"`
	Secret   string                 `json:"-"`
}

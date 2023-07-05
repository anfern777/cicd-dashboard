package entity

import "gorm.io/gorm"

type SourceCodeHostIntegration struct {
	gorm.Model
	ProjectID uint
	Provider  string `validate:"required"`
	Owner     string
	Repo      string
	Secret    string
}

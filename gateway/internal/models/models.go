package models

import "gorm.io/gorm"

type Route struct {
	gorm.Model
	Name        string
	Method      string `gorm:"notnull"`
	Path        string `gorm:"notnull"`
	Enabled     bool   `gorm:"default:true"`
	UpstreamURL string `gorm:"notnull" json:"upstream_url"`
}

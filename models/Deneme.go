package models

import "gorm.io/gorm"

type Deneme struct {
	gorm.Model

	name string `gorm:"type:varchar"`
}

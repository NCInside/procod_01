package models

import "gorm.io/gorm"

type Label struct {
	gorm.Model
	Label string
}

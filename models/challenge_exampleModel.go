package models

import "gorm.io/gorm"

type ChallengeExample struct {
	gorm.Model
	Ex_input    string
	Ex_output   string
	Description string
	ChallengeID uint
}

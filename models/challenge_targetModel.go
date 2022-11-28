package models

import "gorm.io/gorm"

type ChallengeTarget struct {
	gorm.Model
	Input         string
	Target_output string
	ChallengeID   uint
}

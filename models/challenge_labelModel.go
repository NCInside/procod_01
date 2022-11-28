package models

import "gorm.io/gorm"

type ChallengeLabel struct {
	gorm.Model
	Label       string
	ChallengeID uint
}

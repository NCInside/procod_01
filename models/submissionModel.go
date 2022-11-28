package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	Code        string
	Runtime     uint
	Is_correct  bool
	Error       string
	Memory      uint
	UserID      uint
	ChallengeID uint
}

package models

import "gorm.io/gorm"

type Statistic struct {
	gorm.Model
	Total_memory            uint
	Total_runtime           uint
	Num_challenge_attempted uint
	Num_challenge_completed uint
	Num_challenge_made      uint
	UserID                  uint
}

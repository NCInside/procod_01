package models

import "gorm.io/gorm"

type Challenge struct {
	gorm.Model
	Title             string
	Description       string
	ChallengeLabels   []Label `gorm:"many2many:challenge_labels;"`
	ChallengeExamples []ChallengeExample
	ChallengeTargets  []ChallengeTarget
	Submissions       []Submission
	UserID            uint
}

package models

import "gorm.io/gorm"

type Challenge struct {
	gorm.Model
	Title             string
	Description       string
	ChallengeLabels   []ChallengeLabel
	ChallengeExamples []ChallengeExample
	ChallengeTargets  []ChallengeTarget
	Submissions       []Submission
	UserID            uint
}

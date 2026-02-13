package entity

import (
	"time"
)

type PendingUniversity struct {
	Id                 int64      `json:"id" db:"id"`
	Name               string     `json:"name" db:"name"`
	NameEnglish        string     `json:"name_english" db:"name_english"`
	DescriptionEnglish string     `json:"description_english" db:"description_english"`
	City               string     `json:"city" db:"city"`
	Category           string     `json:"category" db:"category"`
	ImageUrl           *string    `json:"image_url,omitempty" db:"image_url"`
	Description        *string    `json:"description,omitempty" db:"description"`
	Status             string     `json:"status" db:"status"`
	SubmittedBy        int64      `json:"submitted_by" db:"submitted_by"`
	SubmittedAt        time.Time  `json:"submitted_at" db:"submitted_at"`
	ApprovedBy         *int64     `json:"approved_by,omitempty" db:"approved_by"`
	ApprovedAt         *time.Time `json:"approved_at,omitempty" db:"approved_at"`
	RejectionReason    *string    `json:"rejection_reason,omitempty" db:"rejection_reason"`
}

type PendingProfessor struct {
	Id                 int64              `json:"id" db:"id"`
	Name               string             `json:"name" db:"name"`
	EducationHistory   *map[string]string `json:"education_history" db:"education_history"` // jsonb
	ImageUrl           *string            `json:"image_url,omitempty" db:"image_url"`
	Description        *string            `json:"description,omitempty" db:"description"`
	Status             string             `json:"status" db:"status"`
	SubmittedBy        int64              `json:"submitted_by" db:"submitted_by"`
	SubmittedAt        time.Time          `json:"submitted_at" db:"submitted_at"`
	NameEnglish        string             `json:"name_english" db:"name_english"`
	DescriptionEnglish string             `json:"description_english" db:"description_english"`
	ApprovedBy         *int64             `json:"approved_by,omitempty" db:"approved_by"`
	ApprovedAt         *time.Time         `json:"approved_at,omitempty" db:"approved_at"`
	RejectionReason    *string            `json:"rejection_reason,omitempty" db:"rejection_reason"`
}

type PendingLesson struct {
	Id                 int64      `json:"id" db:"id"`
	Name               string     `json:"name" db:"name"`
	NameEnglish        string     `json:"name_english" db:"name_english"`
	DescriptionEnglish string     `json:"description_english" db:"description_english"`
	Difficulty         int        `json:"difficulty" db:"difficulty"`
	Description        *string    `json:"description,omitempty" db:"description"`
	Status             string     `json:"status" db:"status"`
	SubmittedBy        int64      `json:"submitted_by" db:"submitted_by"`
	SubmittedAt        time.Time  `json:"submitted_at" db:"submitted_at"`
	ApprovedBy         *int64     `json:"approved_by,omitempty" db:"approved_by"`
	ApprovedAt         *time.Time `json:"approved_at,omitempty" db:"approved_at"`
	RejectionReason    *string    `json:"rejection_reason,omitempty" db:"rejection_reason"`
}
type PendingMajor struct {
	Id                 int64      `json:"id" db:"id"`
	Name               string     `json:"name" db:"name"`
	Status             string     `json:"status" db:"status"`
	NameEnglish        string     `json:"name_english" db:"name_english"`
	SubmittedBy        int64      `json:"submitted_by" db:"submitted_by"`
	Description        *string    `json:"description,omitempty" db:"description"`
	SubmittedAt        time.Time  `json:"submitted_at" db:"submitted_at"`
	DescriptionEnglish string     `json:"description_english" db:"description_english"`
	ApprovedBy         *int64     `json:"approved_by,omitempty" db:"approved_by"`
	ApprovedAt         *time.Time `json:"approved_at,omitempty" db:"approved_at"`
	RejectionReason    *string    `json:"rejection_reason,omitempty" db:"rejection_reason"`
}

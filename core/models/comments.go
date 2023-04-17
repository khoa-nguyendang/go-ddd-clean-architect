package models

type Comment struct {
	BaseModel
	JobId      string  `json:"job_id" db:"job_id,omitempty"`
	UserId     string  `json:"user_id" db:"user_id,omitempty"`
	Content    string  `json:"mid_name" db:"mid_name,omitempty"`
	Rating     float64 `json:"rating" db:"rating,omitempty"`
	IsApproved bool    `json:"is_approved" db:"is_approved,omitempty"`
}

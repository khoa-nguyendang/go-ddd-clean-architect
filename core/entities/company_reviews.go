package entities

type CompanyReview struct {
	BaseEntity
	ParentReviewId      string  `json:"parent_review_id" db:"parent_review_id,omitempty"`
	CompanyId           string  `json:"company_id" db:"company_id,omitempty"`
	UserId              string  `json:"user_id" db:"user_id,omitempty"`
	Content             string  `json:"mid_name" db:"mid_name,omitempty"`
	Rating              float64 `json:"rating" db:"rating,omitempty"`
	IsApproved          bool    `json:"is_approved" db:"is_approved,omitempty"`
	IsAnnonymous        bool    `json:"is_annonymous" db:"is_annonymous,omitempty"`
	Reactions           string  `json:"reactions" db:"reactions,omitempty"`
	Title               string  `json:"title" db:"title,omitempty"`
	WorkingHistoryMonth float64 `json:"working_history_month" db:"working_history_month,omitempty"`
}

package entities

type BaseEntity struct {
	PK              string `json:"pk" db:"pk"`
	Status          string `json:"status" db:"status,omitempty"`
	CreatedDate     string `json:"created_date" db:"created_date"`
	ActivatedDate   string `json:"activated_date" db:"activated_date,omitempty"`
	IsDeleted       bool   `json:"is_deleted" db:"is_deleted,omitempty"`
	DeletedAt       string `json:"deleted_at" db:"deleted_at,omitempty"`
	DeletedByUserId string `json:"deleted_by_user_id" db:"deleted_by_user_id,omitempty"`
}

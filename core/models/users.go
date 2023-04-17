package models

// User reflect user table from authentication service
type User struct {
	BaseModel
	CompanyId      string  `json:"company_name" db:"company_name,omitempty"`
	FirstName      string  `json:"fist_name" db:"fist_name,omitempty"`
	MidName        string  `json:"mid_name" db:"mid_name,omitempty"`
	LastName       string  `json:"last_name" db:"last_name,omitempty"`
	Address        string  `json:"address" db:"address,omitempty"`
	PhoneNumber    string  `json:"phone_number" db:"phone_number,omitempty"`
	Email          string  `json:"email" db:"email,omitempty"`
	UserName       string  `json:"username" db:"username,omitempty"`
	AvatarPath     string  `json:"avatar_path" db:"avatar_path,omitempty"`
	Rating         float64 `json:"rating" db:"rating,omitempty"`
	Title          string  `json:"title" db:"title,omitempty"`
	WorkingMonth   float64 `json:"working_month" db:"working_month,omitempty"`
	CurrentCompany string  `json:"current_company" db:"current_company,omitempty"`
}

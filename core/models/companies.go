package models

type Company struct {
	BaseModel
	CompanyName      string  `json:"company_name" db:"company_name,omitempty"`
	CompanyLegalName string  `json:"company_legal_name" db:"company_legal_name,omitempty"`
	Address          string  `json:"address" db:"address,omitempty"`
	PhoneNumber      string  `json:"phone_number" db:"phone_number,omitempty"`
	TaxId            string  `json:"tax_id" db:"tax_id,omitempty"`
	RegistrationId   string  `json:"registration_id" db:"registration_id,omitempty"`
	ParentCompanyId  string  `json:"parent_company_id" db:"parent_company_id,omitempty"`
	Code             string  `json:"code" db:"code,omitempty"`
	Rating           float64 `json:"rating" db:"rating,omitempty"`
	Type             string  `json:"type" db:"type,omitempty"`
	LogoPath         string  `json:"logo_path" db:"logo_path,omitempty"`
	BackgroundPath   string  `json:"background_path" db:"background_path,omitempty"`
}

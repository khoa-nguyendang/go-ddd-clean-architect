package models

type Job struct {
	BaseModel
	OrderId            string    `json:"order_id" db:"order_id,omitempty"`
	Content            string    `json:"content" db:"content,omitempty"`
	AuthorId           string    `json:"author_id" db:"author_id,omitempty"`
	Title              string    `json:"title" db:"title,omitempty"`
	Tags               string    `json:"tags" db:"tags,omitempty"`
	CompanyId          string    `json:"company_id" db:"company_id,omitempty"`
	Rating             float64   `json:"rating" db:"rating,omitempty"`
	ShortDescription   string    `json:"short_description" db:"short_description,omitempty"`
	FullDescription    string    `json:"full_description" db:"full_description,omitempty"`
	TopicID            string    `json:"topic_id" db:"topic_id,omitempty"`
	Code               string    `json:"code" db:"code,omitempty"`
	ImageThumbnailPath string    `json:"image_thumbnail_path" db:"image_thumbnail_path,omitempty"`
	VideoThumbnailPath string    `json:"video_thumbnail_path" db:"video_thumbnail_path,omitempty"`
	IsApproved         bool      `json:"is_approved" db:"is_approved,omitempty"`
	Comments           []Comment `json:"comments" db:"comments,omitempty"`
	Author             []User    `json:"author" db:"users,omitempty"`
}

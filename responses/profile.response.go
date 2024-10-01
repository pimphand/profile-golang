package responses

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	About string `json:"about"`
}

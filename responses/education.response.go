package responses

import "gorm.io/gorm"

type Education struct {
	gorm.Model
	Degree      string `json:"degree"`
	School      string `json:"school"`
	StartYear   string `json:"start_year"`
	EndYear     string `json:"end_year"`
	Description string `json:"description"`
	Departement string `json:"departement"`
}

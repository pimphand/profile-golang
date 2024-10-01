package responses

type Experience struct {
	ID          uint   `json:"id"`
	Company     string `json:"company"`
	Title       string `json:"title"`
	StartYear   string `json:"start_year"`
	EndYear     string `json:"end_year"`
	Description string `json:"description"`
}

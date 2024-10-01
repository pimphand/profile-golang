package responses

type Project struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Url         string   `json:"url"`
	Image       string   `json:"image"`
	Status      string   `json:"status"`
	Tags        []string `json:"tags"`
}

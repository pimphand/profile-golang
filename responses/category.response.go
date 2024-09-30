package responses

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ParentID   uint64 `json:"parent_id"`
	Name       string `json:"name"`
	IconURL    string `json:"icon_url"`
	IsRootNode bool   `json:"is_root_node"`
}

package responses

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID              uint64   `json:"id"`
	Name            string   `json:"name"`
	CategoryId      uint64   `json:"category_id"`
	Category        Category `gorm:"references:ID" json:"category"` // Define relationship
	Slug            string   `json:"slug"`
	IsRecomendation bool     `json:"is_recomendation"`
	IsLast          bool     `json:"is_last"`
	Skus            []Sku    `gorm:"foreignKey:ProductID" json:"skus"` // One-to-many relationship
}

func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name LIKE ?", "%"+search+"%").Or("slug LIKE ?", search)
	}
}

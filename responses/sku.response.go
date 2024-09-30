package responses

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

// SKU struct represents the SKU entity in the database
type Sku struct {
	gorm.Model                            // Includes ID, CreatedAt, UpdatedAt, and DeletedAt fields
	Code                   string         `json:"code" gorm:"type:varchar(255)"`                  // SKU code
	ProductID              uint64         `json:"product_id" gorm:"not null"`                     // Foreign key to the product
	Properties             PropertiesJSON `json:"properties" gorm:"type:json"`                    // JSON properties
	Stock                  uint64         `json:"stock" gorm:"type:int unsigned"`                 // Stock amount
	WeightGram             uint64         `json:"weight_gram" gorm:"type:int unsigned"`           // Weight in grams
	ConsumerPriceIdr       uint64         `json:"consumer_price_idr" gorm:"type:bigint unsigned"` // Consumer price in IDR
	PriceType              *int           `json:"price_type" gorm:"type:int"`                     // Nullable price type
	NgorderSKUId           *uint64        `json:"ngorder_sku_id" gorm:"type:bigint unsigned"`     // Nullable
	NgorderStock           *uint64        `json:"ngorder_stock" gorm:"type:bigint unsigned"`      // Nullable
	TiktokSKUId            *string        `json:"tiktok_sku_id" gorm:"type:varchar(255)"`         // Nullable
	TiktokStock            *uint64        `json:"tiktok_stock" gorm:"type:bigint unsigned"`       // Nullable
	LazadaSKUId            *string        `json:"lazada_sku_id" gorm:"type:varchar(255)"`         // Nullable
	LazadaStock            *uint64        `json:"lazada_stock" gorm:"type:bigint unsigned"`       // Nullable
	TokopediaSKUId         *string        `json:"tokopedia_sku_id" gorm:"type:varchar(255)"`      // Nullable
	TokopediaStock         *uint64        `json:"tokopedia_stock" gorm:"type:bigint unsigned"`    // Nullable
	IdejualanSKUId         *string        `json:"idejualan_sku_id" gorm:"type:varchar(255)"`      // Nullable
	IdejualanStock         *uint64        `json:"idejualan_stock" gorm:"type:bigint unsigned"`    // Nullable
	ShopeeSKUId            *string        `json:"shopee_sku_id" gorm:"type:varchar(255)"`         // Nullable
	ShopeeStock            *uint64        `json:"shopee_stock" gorm:"type:bigint unsigned"`       // Nullable
	HPP                    *int           `json:"hpp"`                                            // Nullable HPP
	StockCustomer          *int           `json:"stock_customer"`                                 // Nullable
	StockReseler           *int           `json:"stock_reseler"`                                  // Nullable
	StockDistributor       *int           `json:"stock_distributor"`                              // Nullable
	StockResellerGold      *int           `json:"stock_reseller_gold"`                            // Nullable
	StockAgen              *int           `json:"stock_agen"`                                     // Nullable
	StockMember            *int           `json:"stock_member"`                                   // Nullable
	StockDistributorPlus   *int           `json:"stock_distributor_plus"`                         // Nullable
	StockCustomerAffiliate *int           `json:"stock_customer_affiliate"`                       // Nullable
	IsPreorder             *int           `json:"is_preorder"`                                    // Nullable
	DayPreorder            *int           `json:"day_preorder"`                                   // Nullable
	PreorderStart          *time.Time     `json:"preorder_start"`                                 // Nullable
	PreorderEnd            *time.Time     `json:"preorder_end"`                                   // Nullable
	SpecialPrice           *int           `json:"special_price"`                                  // Nullable
	PriceWeb               *int           `json:"price_web"`                                      // Nullable
	Images                 *[]Image       `gorm:"foreignKey:ImageableID;references:ID" json:"images"`
}

// change properties to array of struct
type Properties struct {
	Color    string `json:"color"`
	Size     string `json:"size"`
	Material string `json:"material"`
}

type PropertiesJSON Properties

// Scan implements the sql.Scanner interface
func (p *PropertiesJSON) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan PropertiesJSON")
	}
	return json.Unmarshal(bytes, (*Properties)(p))
}

// Value implements the driver.Valuer interface
func (p PropertiesJSON) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// get Images
type Image struct {
	gorm.Model
	Image         string `json:"image"`
	ImageableID   uint64 `json:"imageable_id"`
	ImageableType string `json:"imageable_type"`
	Path          string `json:"path"`
}

// GetPublicUrl replaces "storage" with "public" in the base URL
func (img *Image) GetPublicUrl() string {
	baseUrl := "https://sgp1.vultrobjects.com/naisha-s3/"
	// Construct the URL
	return baseUrl + img.Path
}

func (img *Image) PublicUrl() string {
	return "https://sgp1.vultrobjects.com/naisha-s3/" + strings.ReplaceAll(img.Path, "storage", "public")
}

// MarshalJSON to customize JSON output for the Image struct
func (img Image) MarshalJSON() ([]byte, error) {
	type Alias Image // Create an alias to avoid recursion in MarshalJSON
	return json.Marshal(&struct {
		PublicUrl string `json:"url"`
		*Alias
	}{
		PublicUrl: img.PublicUrl(),
		Alias:     (*Alias)(&img),
	})
}

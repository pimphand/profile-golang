package models

type Product struct {
	ID                 uint64  `json:"id"`
	SupplierID         uint64  `json:"supplier_id"`
	Name               string  `json:"name"`
	Slug               string  `json:"slug"`
	Description        string  `json:"description,omitempty"`
	Tags               string  `json:"tags"`
	CategoryID         *uint64 `json:"category_id,omitempty"` // Nullable
	Published          bool    `json:"published"`
	ShowStock          bool    `json:"show_stock"`
	CommissionValue    uint64  `json:"commission_value"`
	CommissionType     uint8   `json:"commission_type"`
	RefPrice           uint64  `json:"ref_price"`
	TiktokProductID    *string `json:"tiktok_product_id,omitempty"`    // Nullable
	LazadaProductID    *uint64 `json:"lazada_product_id,omitempty"`    // Nullable
	TokopediaProductID *uint64 `json:"tokopedia_product_id,omitempty"` // Nullable
	ShopeeProductID    *string `json:"shopee_product_id,omitempty"`    // Nullable
	IdejualanProductID *string `json:"idejualan_product_id,omitempty"` // Nullable
	NgorderProductID   *uint64 `json:"ngorder_product_id,omitempty"`   // Nullable
	CreatedAt          *string `json:"created_at,omitempty"`           // Nullable timestamp (use appropriate Go time package)
	UpdatedAt          *string `json:"updated_at,omitempty"`           // Nullable timestamp
	DeletedAt          *string `json:"deleted_at,omitempty"`           // Nullable timestamp
	Cover              *string `json:"cover,omitempty"`                // Nullable
	IsRecomendation    bool    `json:"is_recomendation"`
	TotalOrder         *int    `json:"total_order,omitempty"` // Nullable, default 0
	IsLast             bool    `json:"is_last"`
}

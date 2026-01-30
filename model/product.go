package model

// Product represents a product entity.
type Product struct {
	ID          int64     `json:"id"`
	UUID        string    `json:"uuid"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Price       *float64  `json:"price"`
	Category    *Category `json:"category"`
}

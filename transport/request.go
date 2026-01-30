package transport

// CategoryRequest represents the payload for creating or updating a category.
type CategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PackageRequest represents the payload for creating or updating a package.
type PackageRequest struct {
	Name        string   `json:"name"`
	CategoryID  string   `json:"category_id"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
}

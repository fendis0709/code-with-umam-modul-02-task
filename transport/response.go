package transport

// StatusResponse represents a standard status response.
type StatusResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

// ProductsResponse represents the response for multiple products.
type ProductsResponse struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    ProductsResponseData `json:"data"`
}

// ProductsResponseData holds the list of products in the response.
type ProductsResponseData struct {
	Products []ProductItemResponse `json:"products"`
}

// ProductResponse represents the response for a single product.
type ProductResponse struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    ProductResponseData `json:"data"`
}

// ProductResponseData holds a single product in the response.
type ProductResponseData struct {
	Product ProductItemResponse `json:"product"`
}

// ProductItemResponse represents a product item in the response.
type ProductItemResponse struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Description *string               `json:"description"`
	Price       *float64              `json:"price"`
	Category    *CategoryItemResponse `json:"category"`
}

// CategoryItemResponse represents a category item in the response.
type CategoryItemResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

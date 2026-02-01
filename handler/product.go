package handler

import (
	"encoding/json"
	"fendi/modul-02-task/service"
	"net/http"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) HandleProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.GetAllProduct(w, r)
		return
	}

	http.NotFound(w, r)
}

func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAllProduct(r.Context())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) FindProductByID(w http.ResponseWriter, r *http.Request) {}

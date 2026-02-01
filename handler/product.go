package handler

import (
	"encoding/json"
	"fendi/modul-02-task/service"
	"fendi/modul-02-task/transport"
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
	if r.Method == http.MethodPost {
		h.CreateProduct(w, r)
		return
	}

	http.NotFound(w, r)
}

func (h *ProductHandler) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetAllProduct(r.Context())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productReq transport.PackageRequest
	err := json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	res, err := h.service.CreateProduct(r.Context(), productReq)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	json.NewEncoder(w).Encode(res)
}

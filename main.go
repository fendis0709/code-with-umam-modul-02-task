package main

import (
	"encoding/json"
	"fendi/modul-02-task/config"
	"fendi/modul-02-task/database"
	"fendi/modul-02-task/handler"
	"fendi/modul-02-task/repository"
	"fendi/modul-02-task/service"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type StatusResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type CategoriesResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    CategoriesResponseData `json:"data"`
}

type CategoriesResponseData struct {
	Categories []Category `json:"categories"`
}

type CategoryResponse struct {
	Code   int                  `json:"code"`
	Status string               `json:"status"`
	Data   CategoryResponseData `json:"data"`
}

type CategoryResponseData struct {
	Category Category `json:"category"`
}

var categories = []Category{
	{
		ID:          "c231e125",
		Name:        "Food",
		Description: "Food and beverages",
	},
}

func main() {
	var err error

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_, err = os.Stat(".env")
	if err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	conf := config.Config{
		AppPort: viper.GetString("APP_PORT"),
		DBConn:  viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(conf.DBConn)
	if err != nil {
		log.Fatalf("Error: Unable to connect to database, %v", err.Error())
	}
	defer db.Close()

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if r.URL.Path == "/" {
				w.Header().Set("Content-Type", "application/json")

				json.NewEncoder(w).Encode(StatusResponse{
					Code:   200,
					Status: "OK",
				})
				return
			}
		}

		http.Error(w, "Not Found", http.StatusNotFound)
	})

	http.HandleFunc("/products", productHandler.HandleProduct)

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategories(w, r)
			return
		}

		if r.Method == "POST" {
			createCategory(w, r)
			return
		}

		http.Error(w, "Not Found", http.StatusNotFound)
	})

	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getDetailCategory(w, r)
			return
		}

		if r.Method == "PUT" {
			updateCategory(w, r)
			return
		}

		if r.Method == "DELETE" {
			deleteCategory(w, r)
			return
		}

		http.Error(w, "Not Found", http.StatusNotFound)
	})

	fmt.Println("Server is up and running")
	fmt.Printf("http://localhost:%s\n", conf.AppPort)

	addr := fmt.Sprintf(":%s", conf.AppPort)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Error: Unable to start server, %v", err.Error())
		return
	}
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(CategoriesResponse{
		Code:    200,
		Message: "OK",
		Data: CategoriesResponseData{
			Categories: categories,
		},
	})
}

func getDetailCategory(w http.ResponseWriter, r *http.Request) {
	paramID := strings.TrimPrefix(r.URL.Path, "/categories/")
	if paramID == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	var category Category
	for _, c := range categories {
		if c.ID == paramID {
			category = c
			break
		}
	}

	if category.ID == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(CategoryResponse{
		Code:   200,
		Status: "OK",
		Data: CategoryResponseData{
			Category: category,
		},
	})
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	newCategory.ID = randomHexaDecimal(8)
	categories = append(categories, newCategory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	json.NewEncoder(w).Encode(CategoryResponse{
		Code:   201,
		Status: "Created",
		Data: CategoryResponseData{
			Category: newCategory,
		},
	})
}

func updateCategory(w http.ResponseWriter, r *http.Request) {
	paramID := strings.TrimPrefix(r.URL.Path, "/categories/")
	if paramID == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	var bodyCategory Category
	err := json.NewDecoder(r.Body).Decode(&bodyCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	bodyCategory.ID = paramID

	var isFound bool
	for i := range categories {
		if categories[i].ID == paramID {
			categories[i] = bodyCategory
			isFound = true
		}
	}

	if !isFound {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(StatusResponse{
		Code:   200,
		Status: "OK",
	})
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	paramID := strings.TrimPrefix(r.URL.Path, "/categories/")
	if paramID == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	newCategories := []Category{}
	var isFound bool
	for _, c := range categories {
		if c.ID == paramID {
			isFound = true
			continue
		}
		newCategories = append(newCategories, c)
	}

	if !isFound {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	categories = newCategories

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(StatusResponse{
		Code:   200,
		Status: "OK",
	})
}

func randomHexaDecimal(n int) string {
	const letters = "abcdef0123456789"

	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

package controllers

import (
	"fmt"
	"go-jwt/configs"
	"go-jwt/helpers"
	"go-jwt/models"
	"net/http"
	"path/filepath"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	var product []models.Product

	db := configs.DB
	result := db.Find(&product)

	if result.Error != nil {
		helpers.Response(w, 500, result.Error.Error(), nil)
		return
	}

	for i := range product {
		product[i].Image = fmt.Sprintf("http://localhost:8080/%s", filepath.ToSlash(product[i].Image))
	}

	helpers.Response(w, 200, "Product created successfully", product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// var product models.Product

	// db := configs.DB
	// var product models.Product
	// json.NewDecoder(r.Body).Decode(&product)
	// res := db.Create(&product)
	// if res.Error != nil {
	// 	helpers.Response(w, 500, res.Error.Error(), nil)
	// }

	// helpers.Response(w, 200, "Product created successfully", product)
	// if err != nil {
	// 	http.Error(w, "Error parsing form data", http.StatusBadRequest)
	// 	return
	// }

	// Access the form data
	r.ParseMultipartForm(32 << 20)

	var product models.Product

	db := configs.DB

	product.Name = r.FormValue("name") // Works for both form-urlencoded and multipart/form-data
	if product.Name == "" {
		http.Error(w, "Name field is required", http.StatusBadRequest)
		return
	}

	product.Price = r.FormValue("price") // Works for both form-urlencoded and multipart/form-data
	if product.Price == "" {
		http.Error(w, "Price field is required", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	uploadDir := "./uploads"
	filePath, err := helpers.SaveFile(file, handler, uploadDir)
	product.Image = "/" + filePath
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save image: %v", err), http.StatusInternalServerError)
		return
	}

	db.Select("Name", "Price", "Image").Create(&product)

	// response := fmt.Sprintf("File '%s' uploaded successfully with name: '%s'", filepath.Base(filePath), name)
	// Send a response
	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Product created successfully"))

	helpers.Response(w, 200, "Product created successfully", nil)
}

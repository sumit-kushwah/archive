package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"your-app/models"
	"your-app/services"
)

type ProductHandler struct {
	Service *services.ProductService
}

// GetProductsHandler retrieves all products
func (ph *ProductHandler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	products, err := ph.Service.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetProductByIDHandler retrieves a specific product by its ID
func (ph *ProductHandler) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.URL.Query().Get("product_id"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := ph.Service.GetProductByID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// CreateProductHandler creates a new product
func (ph *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate and process the new product
	createdProduct, err := ph.Service.CreateProduct(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(createdProduct)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// UpdateProductHandler updates an existing product by its ID
func (ph *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.URL.Query().Get("product_id"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct models.Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate and process the updated product
	updatedProduct, err = ph.Service.UpdateProduct(productID, &updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(updatedProduct)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// DeleteProductHandler deletes a product by its ID
func (ph *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.URL.Query().Get("product_id"))
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = ph.Service.DeleteProduct(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetProductsByCategoryHandler retrieves all products within a specific category
func (ph *ProductHandler) GetProductsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.Atoi(r.URL.Query().Get("category_id"))
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	products, err := ph.Service.GetProductsByCategory(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

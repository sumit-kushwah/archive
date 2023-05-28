package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"your-app/models"
	"your-app/services"
)

type ProductCategoryHandler struct {
	Service *services.ProductCategoryService
}

// GetProductCategoriesHandler retrieves all product categories
func (ph *ProductCategoryHandler) GetProductCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	productCategories, err := ph.Service.GetAllProductCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(productCategories)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// GetProductCategoryByIDHandler retrieves a specific product category by its ID
func (ph *ProductCategoryHandler) GetProductCategoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.Atoi(r.URL.Query().Get("category_id"))
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	productCategory, err := ph.Service.GetProductCategoryByID(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(productCategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// CreateProductCategoryHandler creates a new product category
func (ph *ProductCategoryHandler) CreateProductCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var newProductCategory models.ProductCategory
	err := json.NewDecoder(r.Body).Decode(&newProductCategory)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate and process the new product category
	createdProductCategory, err := ph.Service.CreateProductCategory(&newProductCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(createdProductCategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

// UpdateProductCategoryHandler updates an existing product category by its ID
func (ph *ProductCategoryHandler) UpdateProductCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.Atoi(r.URL.Query().Get("category_id"))
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var updatedProductCategory models.ProductCategory
	err = json.NewDecoder(r.Body).Decode(&updatedProductCategory)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate and process the updated product category
	updatedProductCategory, err = ph.Service.UpdateProductCategory(categoryID, &updatedProductCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(updatedProductCategory)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// DeleteProductCategoryHandler deletes a product category by its ID
func (ph *ProductCategoryHandler) DeleteProductCategoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryID, err := strconv.Atoi(r.URL.Query().Get("category_id"))
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	err = ph.Service.DeleteProductCategory(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

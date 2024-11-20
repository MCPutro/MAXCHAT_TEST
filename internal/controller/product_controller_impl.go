package controller

import (
	"encoding/json"
	"fmt"
	"github.com/MCPutro/maxchatTest/internal/domain"
	"github.com/MCPutro/maxchatTest/internal/usecase"
	"github.com/MCPutro/maxchatTest/pkg/utils"
	"net/http"
)

// ProductHandler defines the structure for handling product-related HTTP requests
type productHandler struct {
	productUseCase usecase.ProductUseCase
}

func NewProductHandler(useCase usecase.ProductUseCase) ProductController {
	return &productHandler{productUseCase: useCase}
}

//// SeedProductsFromJSON handles the seeding of products from a JSON file
//func (h *productHandler) SeedProductsFromJSON(w http.ResponseWriter, r *http.Request) {
//	// Read file path from query parameters
//	filePath := r.URL.Query().Get("file")
//	if filePath == "" {
//		http.Error(w, "Missing file path", http.StatusBadRequest)
//		return
//	}
//
//	// Seed products from the JSON file
//	if err := h.usecase.SeedProductsFromJSON(filePath); err != nil {
//		http.Error(w, fmt.Sprintf("Error seeding products: %v", err), http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte("Products seeded successfully"))
//}

func (h *productHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	allProducts, err := h.productUseCase.ReadAllProducts()
	if err != nil {
		http.Error(w, "Missing tech parameter", http.StatusBadRequest)
		return
	}

	// Respond with products as JSON
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(allProducts)
	utils.WriteToResponseBody(err, fmt.Sprintf("Error retrieving products : %v", err), w, allProducts)
}

// GetProductsByTech handles the retrieval of products by tech
func (h *productHandler) GetProductsByTech(w http.ResponseWriter, r *http.Request) {
	//tech := r.URL.Query().Get("tech")
	tech := r.PathValue("tech")
	if tech == "" {
		http.Error(w, "Missing tech parameter", http.StatusBadRequest)
		return
	}

	products, err := h.productUseCase.ReadProductsByTech(tech)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Error retrieving products by tech: %v", err), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(products)

	utils.WriteToResponseBody(err, fmt.Sprintf("Error retrieving products by tech: %v", err), w, products)
}

// GetProductsByModel handles the retrieval of products by model
func (h *productHandler) GetProductsByModel(w http.ResponseWriter, r *http.Request) {
	//model := r.URL.Query().Get("model")
	model := r.PathValue("model")
	if model == "" {
		http.Error(w, "Missing model parameter", http.StatusBadRequest)
		return
	}

	products, err := h.productUseCase.ReadProductsByModel(model)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Error retrieving products by model: %v", err), http.StatusInternalServerError)
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(products)

	utils.WriteToResponseBody(err, fmt.Sprintf("Error retrieving products by model: %v", err), w, products)
}

func (h *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.productUseCase.UpdateProduct(&product)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Error updating product: %v", err), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Product updated successfully"))

	utils.WriteToResponseBody(err, fmt.Sprintf("Error updating product: %v", err), w, nil)
}

// DeleteProduct handles the deletion of a product
func (h *productHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	//code := r.URL.Query().Get("code")
	code := r.PathValue("code")
	if code == "" {
		http.Error(w, "Missing product code", http.StatusBadRequest)
		return
	}

	err := h.productUseCase.DeleteProduct(code)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Error deleting product: %v", err), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Product deleted successfully"))

	utils.WriteToResponseBody(err, fmt.Sprintf("Error deleting product: %v", err), w, nil)
}

func (h *productHandler) AddNewProduct(w http.ResponseWriter, r *http.Request) {
	var request domain.Product
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.productUseCase.AddProduct(&request)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("Error add product: %v", err), http.StatusInternalServerError)
	//	return
	//}
	//
	//w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Add product successfully"))

	utils.WriteToResponseBody(err, fmt.Sprintf("EError add product: %v", err), w, nil)

}

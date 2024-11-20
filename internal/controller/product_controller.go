package controller

import (
	"net/http"
)

type ProductController interface {
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	GetProductsByTech(w http.ResponseWriter, r *http.Request)
	GetProductsByModel(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	AddNewProduct(w http.ResponseWriter, r *http.Request)
}

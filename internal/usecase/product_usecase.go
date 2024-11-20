package usecase

import (
	"github.com/MCPutro/maxchatTest/internal/domain"
)

type ProductUseCase interface {
	SeedProductsFromJSON(filePath string) error
	ReadProductsByTech(tech string) ([]*domain.Product, error)
	ReadProductsByModel(model string) ([]*domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(code string) error
	ReadProductsByCode(code string) (*domain.Product, error)
	ReadAllProducts() ([]*domain.Product, error)
	AddProduct(product *domain.Product) error
}

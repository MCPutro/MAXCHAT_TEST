package repository

import "github.com/MCPutro/maxchatTest/internal/domain"

type ProductRepository interface {
	Write(product *domain.Product) error
	Read(code string) (*domain.Product, error)
	ReadAll() ([]*domain.Product, error)
	ReadByTech(tech string) ([]*domain.Product, error)
	ReadByModel(model string) ([]*domain.Product, error)
	Update(product *domain.Product) error
	Delete(code string) error
}

package usecase

import (
	"encoding/json"
	"errors"
	"github.com/MCPutro/maxchatTest/internal/domain"
	"github.com/MCPutro/maxchatTest/internal/repository"
	"os"
)

type productUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}

func (u *productUseCase) SeedProductsFromJSON(filePath string) error {
	//file, err := os.Open(filePath)
	//if err != nil {
	//	return errors.New("failed to open JSON file: " + err.Error())
	//}
	//defer file.Close()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return errors.New("failed to read JSON file: " + err.Error())
	}

	var products []*domain.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return errors.New("failed to parse JSON: " + err.Error())
	}

	for _, product := range products {
		if err := u.repo.Write(product); err != nil {
			return errors.New("failed to seed product " + product.Code + ": " + err.Error())
		}
	}

	return nil
}

func (u *productUseCase) ReadProductsByTech(tech string) ([]*domain.Product, error) {
	return u.repo.ReadByTech(tech)
}

func (u *productUseCase) ReadProductsByModel(model string) ([]*domain.Product, error) {
	return u.repo.ReadByModel(model)
}

func (u *productUseCase) UpdateProduct(product *domain.Product) error {
	return u.repo.Update(product)
}

func (u *productUseCase) DeleteProduct(code string) error {
	return u.repo.Delete(code)
}

func (u *productUseCase) ReadProductsByCode(code string) (*domain.Product, error) {
	return u.repo.Read(code)
}

func (u *productUseCase) ReadAllProducts() ([]*domain.Product, error) {
	return u.repo.ReadAll()
}

func (u *productUseCase) AddProduct(product *domain.Product) error {
	return u.repo.Write(product)
}

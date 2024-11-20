package repository

import (
	"errors"
	"fmt"
	"sync"

	"github.com/MCPutro/maxchatTest/internal/domain"
)

var (
	products        map[string]*domain.Product
	productOnce     sync.Once
	productsByTech  map[string][]string
	productsByModel map[string][]string
)

type ProductRepo struct {
	mu sync.Mutex
}

//func initProducts() {
//	products = make(map[string]*domain.Product)
//	productsByTech = make(map[string][]string)
//	productsByModel = make(map[string][]string)
//}

func NewProductRepo() ProductRepository {
	productOnce.Do(func() {
		products = make(map[string]*domain.Product)
		productsByTech = make(map[string][]string)
		productsByModel = make(map[string][]string)
	})
	return &ProductRepo{}
}

func (repo *ProductRepo) Write(product *domain.Product) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if _, exists := products[product.Code]; exists {
		return errors.New(fmt.Sprintf("product with code %s already exists", product.Code))
	}

	products[product.Code] = product

	for _, tech := range product.Tech {
		productsByTech[tech] = append(productsByTech[tech], product.Code)
	}
	productsByModel[product.Model] = append(productsByModel[product.Model], product.Code)

	return nil
}

func (repo *ProductRepo) Read(code string) (*domain.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	product, exists := products[code]
	if !exists {
		return nil, errors.New("product not found")
	}

	return product, nil
}

func (repo *ProductRepo) ReadAll() ([]*domain.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	var result []*domain.Product
	for _, p := range products {
		result = append(result, p)
	}

	return result, nil
}

func (repo *ProductRepo) ReadByTech(tech string) ([]*domain.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	ids, exists := productsByTech[tech]
	if !exists || len(ids) == 0 {
		return nil, errors.New(fmt.Sprintf("no products found for tech %s", tech))
	}

	var result []*domain.Product
	for _, code := range ids {
		result = append(result, products[code])
	}
	return result, nil
}

func (repo *ProductRepo) ReadByModel(model string) ([]*domain.Product, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	ids, exists := productsByModel[model]
	if !exists || len(ids) == 0 {
		return nil, errors.New(fmt.Sprintf("no products found for model %s", model))
	}

	var result []*domain.Product
	for _, code := range ids {
		result = append(result, products[code])
	}
	return result, nil
}

func (repo *ProductRepo) Update(product *domain.Product) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	existingProduct, exists := products[product.Code]
	if !exists {
		return errors.New("product not found")
	}

	*existingProduct = *product

	for _, tech := range existingProduct.Tech {
		productsByTech[tech] = removeProductFromList(productsByTech[tech], product.Code)
	}

	productsByModel[existingProduct.Model] = removeProductFromList(productsByModel[existingProduct.Model], product.Code)

	for _, tech := range product.Tech {
		productsByTech[tech] = append(productsByTech[tech], product.Code)
	}
	productsByModel[product.Model] = append(productsByModel[product.Model], product.Code)

	return nil
}

func (repo *ProductRepo) Delete(code string) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	// Check if the product exists
	product, exists := products[code]
	if !exists {
		return errors.New("product not found")
	}

	// Delete the product from the map
	delete(products, code)

	// Remove from productsByTech and productsByModel
	for _, tech := range product.Tech {
		productsByTech[tech] = removeProductFromList(productsByTech[tech], code)
	}
	productsByModel[product.Model] = removeProductFromList(productsByModel[product.Model], code)

	return nil
}

func removeProductFromList(list []string, code string) []string {
	for i, item := range list {
		if item == code {
			list = append(list[:i], list[i+1:]...)
			break
		}
	}
	return list
}

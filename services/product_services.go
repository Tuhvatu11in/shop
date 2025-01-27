package services

import (
	"shop/models"
	"sync"
)

type ProductService struct {
	products map[int]models.Product
	mu       sync.RWMutex
	nextID   int
}

func NewProductService() *ProductService {
	return &ProductService{
		products: make(map[int]models.Product),
		nextID:   1,
	}
}

func (ps *ProductService) AddProduct(p models.Product) int {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	p.ID = ps.nextID
	ps.products[ps.nextID] = p
	ps.nextID++
	return p.ID
}

func (ps *ProductService) GetProduct(id int) (models.Product, bool) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	p, ok := ps.products[id]
	return p, ok
}

func (ps *ProductService) GetAllProducts() []models.Product {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	var products []models.Product
	for _, p := range ps.products {
		products = append(products, p)
	}
	return products
}

func (ps *ProductService) UpdateProduct(id int, p models.Product) bool {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if _, ok := ps.products[id]; ok {
		p.ID = id
		ps.products[id] = p
		return true
	}
	return false
}

func (ps *ProductService) DeleteProduct(id int) bool {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if _, ok := ps.products[id]; ok {
		delete(ps.products, id)
		return true
	}
	return false
}

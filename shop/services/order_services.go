package services

import (
	"shop/models"
	"sync"
)

type OrderService struct {
	orders         map[int]models.Order
	mu             sync.RWMutex
	nextID         int
	productService *ProductService
}

func NewOrderService(ps *ProductService) *OrderService {
	return &OrderService{
		orders:         make(map[int]models.Order),
		nextID:         1,
		productService: ps,
	}
}

func (os *OrderService) AddOrder(o models.Order) int {
	os.mu.Lock()
	defer os.mu.Unlock()
	o.ID = os.nextID
	total := 0.0
	for _, item := range o.Items {
		product, ok := os.productService.GetProduct(item.ProductID)
		if ok {
			total += product.Price * float64(item.Quantity)
		}
	}
	o.Total = total
	os.orders[os.nextID] = o
	os.nextID++
	return o.ID
}

func (os *OrderService) GetOrder(id int) (models.Order, bool) {
	os.mu.RLock()
	defer os.mu.RUnlock()
	o, ok := os.orders[id]
	return o, ok
}

func (os *OrderService) GetAllOrders() []models.Order {
	os.mu.RLock()
	defer os.mu.RUnlock()
	var orders []models.Order
	for _, o := range os.orders {
		orders = append(orders, o)
	}
	return orders
}

func (os *OrderService) UpdateOrder(id int, o models.Order) bool {
	os.mu.Lock()
	defer os.mu.Unlock()
	if _, ok := os.orders[id]; ok {
		o.ID = id
		total := 0.0
		for _, item := range o.Items {
			product, ok := os.productService.GetProduct(item.ProductID)
			if ok {
				total += product.Price * float64(item.Quantity)
			}
		}
		o.Total = total
		os.orders[id] = o
		return true
	}
	return false
}

func (os *OrderService) DeleteOrder(id int) bool {
	os.mu.Lock()
	defer os.mu.Unlock()
	if _, ok := os.orders[id]; ok {
		delete(os.orders, id)
		return true
	}
	return false
}

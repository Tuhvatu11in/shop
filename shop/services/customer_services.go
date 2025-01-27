package services

import (
	"shop/models"
	"sync"
)

type CustomerService struct {
	customers map[int]models.Customer
	mu        sync.RWMutex
	nextID    int
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		customers: make(map[int]models.Customer),
		nextID:    1,
	}
}
func (cs *CustomerService) AddCustomer(c models.Customer) int {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	c.ID = cs.nextID
	cs.customers[cs.nextID] = c
	cs.nextID++
	return c.ID
}

func (cs *CustomerService) GetCustomer(id int) (models.Customer, bool) {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	c, ok := cs.customers[id]
	return c, ok
}
func (cs *CustomerService) GetAllCustomers() []models.Customer {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	var customers []models.Customer
	for _, c := range cs.customers {
		customers = append(customers, c)
	}
	return customers
}

func (cs *CustomerService) UpdateCustomer(id int, c models.Customer) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	if _, ok := cs.customers[id]; ok {
		c.ID = id
		cs.customers[id] = c
		return true
	}
	return false
}

func (cs *CustomerService) DeleteCustomer(id int) bool {
	cs.mu.Lock()
	defer cs.mu.Unlock()
	if _, ok := cs.customers[id]; ok {
		delete(cs.customers, id)
		return true
	}
	return false
}

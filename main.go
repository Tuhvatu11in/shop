package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"shop/models"
	"shop/services"
	"shop/utils"
)

var (
	productService  *services.ProductService
	customerService *services.CustomerService
	orderService    *services.OrderService
)

func main() {
	productService = services.NewProductService()
	customerService = services.NewCustomerService()
	orderService = services.NewOrderService(productService)

	r := mux.NewRouter()

	r.HandleFunc("/products", handleProducts).Methods("GET", "POST")
	r.HandleFunc("/products/{id}", handleProduct).Methods("GET", "PUT", "DELETE")

	r.HandleFunc("/customers", handleCustomers).Methods("GET", "POST")
	r.HandleFunc("/customers/{id}", handleCustomer).Methods("GET", "PUT", "DELETE")

	r.HandleFunc("/orders", handleOrders).Methods("GET", "POST")
	r.HandleFunc("/orders/{id}", handleOrder).Methods("GET", "PUT", "DELETE")

	// Запуск сервера
	corsHandler := cors.AllowAll().Handler(r)
	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		products := productService.GetAllProducts()
		utils.RespondWithJSON(w, http.StatusOK, products)
	case http.MethodPost:
		var newProduct models.Product
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		id := productService.AddProduct(newProduct)
		utils.RespondWithJSON(w, http.StatusCreated, map[string]int{"id": id})
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func handleProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	switch r.Method {
	case http.MethodGet:
		product, ok := productService.GetProduct(id)
		if !ok {
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, product)
	case http.MethodPut:
		var updatedProduct models.Product
		err := json.NewDecoder(r.Body).Decode(&updatedProduct)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if !productService.UpdateProduct(id, updatedProduct) {
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Product updated"})
	case http.MethodDelete:
		if !productService.DeleteProduct(id) {
			utils.RespondWithError(w, http.StatusNotFound, "Product not found")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Product deleted"})
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func handleCustomers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		customers := customerService.GetAllCustomers()
		utils.RespondWithJSON(w, http.StatusOK, customers)
	case http.MethodPost:
		var newCustomer models.Customer
		err := json.NewDecoder(r.Body).Decode(&newCustomer)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		id := customerService.AddCustomer(newCustomer)
		utils.RespondWithJSON(w, http.StatusCreated, map[string]int{"id": id})
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func handleCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	switch r.Method {
	case http.MethodGet:
		customer, ok := customerService.GetCustomer(id)
		if !ok {
			utils.RespondWithError(w, http.StatusNotFound, "Customer not found")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, customer)
	case http.MethodPut:
		var updatedCustomer models.Customer
		err := json.NewDecoder(r.Body).Decode(&updatedCustomer)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if !customerService.UpdateCustomer(id, updatedCustomer) {
			utils.RespondWithError(w, http.StatusNotFound, "Customer not found")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Customer updated"})
	case http.MethodDelete:
		if !customerService.DeleteCustomer(id) {
			utils.RespondWithError(w, http.StatusNotFound, "Customer not found")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Customer deleted"})
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func handleOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		orders := orderService.GetAllOrders()
		utils.RespondWithJSON(w, http.StatusOK, orders)
	case http.MethodPost:
		var newOrder models.Order
		err := json.NewDecoder(r.Body).Decode(&newOrder)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		id := orderService.AddOrder(newOrder)
		utils.RespondWithJSON(w, http.StatusCreated, map[string]int{"id": id})
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	switch r.Method {
	case http.MethodGet:
		order, ok := orderService.GetOrder(id)
		if !ok {
			utils.RespondWithError(w, http.StatusNotFound, "Order not found")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, order)
	case http.MethodPut:
		var updatedOrder models.Order
		err := json.NewDecoder(r.Body).Decode(&updatedOrder)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		if !orderService.UpdateOrder(id, updatedOrder) {
			utils.RespondWithError(w, http.StatusNotFound, "Order not found")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Order updated"})
	case http.MethodDelete:
		if !orderService.DeleteOrder(id) {
			utils.RespondWithError(w, http.StatusNotFound, "Order not found")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Order deleted"})
	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

package interfaces

import "dbclient/models"

type ICustomer interface {
	AddCustomer(customer *models.Customer) (string, error)
	GetCustomers() ([]*models.Customer, error)
}
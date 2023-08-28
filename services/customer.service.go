package services

import (
	"context"
	"dbclient/interfaces"
	"dbclient/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type BankingService struct {
	BankingCollection *mongo.Collection
	ctx               context.Context
}

// GetCustomers implements interfaces.ICustomer.
func (*BankingService) GetCustomers() ([]*models.Customer, error) {
	panic("unimplemented")
}

func NewBankingServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &BankingService{collection, ctx}

}

func (t *BankingService) AddCustomer(customer *models.Customer) (string, error) {

	//hashedPassword, _ := utils.HashPassword(user.Password)

	_, err := t.BankingCollection.InsertOne(t.ctx, &customer)
	if err != nil {
		return "error", nil
	}

	return "success", nil
}

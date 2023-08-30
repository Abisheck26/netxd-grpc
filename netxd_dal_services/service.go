package netxd_dal_services

import (
	"context"
	"time"

	"github.com/Abisheck26/netxd-grpc/netxd_dal_interfaces"
	"github.com/Abisheck26/netxd-grpc/netxd_dal_models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func NewCustomerServiceInit(collection *mongo.Collection, ctx context.Context) netxd_dal_interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (c *CustomerService) CreateCustomer(customer *netxd_dal_models.Customer) (*netxd_dal_models.DBResponse, error) {
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = customer.CreatedAt
	customer.IsActive = true

	res, err := c.CustomerCollection.InsertOne(c.ctx, &customer)

	if err != nil {
		return nil, err
	}

	var newUser *netxd_dal_models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = c.CustomerCollection.FindOne(c.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil

}

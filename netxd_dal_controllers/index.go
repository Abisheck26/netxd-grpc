package controllers

import (
	"context"

	pro "github.com/Abisheck26/netxd-grpc/netxd_customer"
	"github.com/Abisheck26/netxd-grpc/netxd_dal_interfaces"
	"github.com/Abisheck26/netxd-grpc/netxd_dal_models"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService netxd_dal_interfaces.ICustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	dbCustomer := &netxd_dal_models.Customer{CustomerId:req.Customerid,FirstName: req.FirstName,LastName: req.LastName,BankId: req.BankId,Balance: req.Balance,IsActive: req.IsActive}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerId: result.CustomerId,
		}
		return responseCustomer, nil
	}
}

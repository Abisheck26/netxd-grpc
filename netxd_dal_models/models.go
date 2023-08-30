package netxd_dal_models

import (
	"time"
)

type Customer struct {
	CustomerId int32     `json:"customerid" bson:"customerid" binding:"required"`
	FirstName  string    `json:"firstname" bson:"firstname" binding:"required"`
	LastName   string    `json:"lastname" bson:"lastname" binding:"required"`
	BankId     int64     `json:"bankid" bson:"bankid,omitempty" binding:"required"`
	Balance    float32   `json:"balance" bson:"balance"`
	CreatedAt  time.Time `json:"createdat" bson:"createdat"`
	UpdatedAt  time.Time `json:"updatedat" bson:"updatedat"`
	IsActive   bool      `json:"isactive" bson:"isactive"`
}

type DBResponse struct {
	CreatedAt time.Time `json:"createdat" bson:"createdat"`
	CustomerId int32     `json:"customerid" bson:"customerid"`
}

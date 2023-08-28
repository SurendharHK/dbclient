package models

type Customer struct{
	Name      string  `json:"name" bson:"name"`
	Balance float64 `json:"balance" bson:"balance"`
}
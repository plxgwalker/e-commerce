// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type CreateOrderInput struct {
	Products        []string  `json:"products"`
	OrderDate       time.Time `json:"orderDate"`
	ShippingAddress *string   `json:"shippingAddress,omitempty"`
	Status          *string   `json:"status,omitempty"`
	CustomerEmail   string    `json:"customerEmail"`
	PaymentStatus   *string   `json:"paymentStatus,omitempty"`
}

type CreateUserInput struct {
	Name  *string `json:"name,omitempty"`
	Email string  `json:"email"`
	Phone *string `json:"phone,omitempty"`
}

type Mutation struct {
}

type Order struct {
	ID              string    `json:"id"`
	Products        []string  `json:"products"`
	OrderDate       time.Time `json:"orderDate"`
	ShippingAddress *string   `json:"shippingAddress,omitempty"`
	Status          *string   `json:"status,omitempty"`
	CustomerEmail   string    `json:"customerEmail"`
	PaymentStatus   *string   `json:"paymentStatus,omitempty"`
}

type Query struct {
}

type User struct {
	ID     string   `json:"id"`
	Name   *string  `json:"name,omitempty"`
	Email  string   `json:"email"`
	Phone  *string  `json:"phone,omitempty"`
	Orders []*Order `json:"orders,omitempty"`
}

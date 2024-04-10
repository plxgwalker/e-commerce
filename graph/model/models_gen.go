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

type CreateProductInput struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Amount      int      `json:"amount"`
	Price       float64  `json:"price"`
	Images      []string `json:"images"`
	Category    string   `json:"category"`
}

type CreateUserInput struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Phone    *string `json:"phone,omitempty"`
	Password *string `json:"password,omitempty"`
	Image    *string `json:"image,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

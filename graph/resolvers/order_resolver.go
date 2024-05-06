package resolvers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"server.go/constants"
	"server.go/graph/model"
	"server.go/models"
	"server.go/services"
	errors "server.go/utils/errors"
	typesConverters "server.go/utils/types_converters"
)

type OrderResolver struct {
	orderService       *services.OrderService
	orderErrors        errors.OrderErrors
	orderTypeConverter typesConverters.OrderTypesConverter
}

func NewOrderResolver() *OrderResolver {
	return &OrderResolver{
		orderService: services.NewOrderService(),
	}
}

func (r *OrderResolver) ID(ctx context.Context, obj *models.Order) (string, error) {
	return obj.Id.Hex(), nil
}

func (r *OrderResolver) Currency(ctx context.Context, obj *models.Order) (model.Currency, error) {
	var currency model.Currency
	switch obj.Currency {
	case constants.USD:
		currency = model.CurrencyUsd
	case constants.EUR:
		currency = model.CurrencyEur
	default:
		currency = model.CurrencyUah
	}
	return currency, nil
}

func (r *OrderResolver) Category(ctx context.Context, obj *models.Order) ([]model.Category, error) {
	var categories []model.Category
	for _, category := range obj.Category {
		switch category {
		case constants.Electronics:
			categories = append(categories, model.CategoryElectronics)
		case constants.Fashion:
			categories = append(categories, model.CategoryFashion)
		case constants.Home:
			categories = append(categories, model.CategoryHome)
		case constants.Sports:
			categories = append(categories, model.CategorySports)
		case constants.Books:
			categories = append(categories, model.CategoryBooks)
		case constants.Automotive:
			categories = append(categories, model.CategoryAutomotive)
		case constants.Other:
			categories = append(categories, model.CategoryOther)
		}
	}
	return categories, nil
}

func (r *OrderResolver) Status(ctx context.Context, obj *models.Order) (model.Status, error) {
	var status model.Status
	switch obj.Status {
	case constants.Available:
		status = model.StatusAvailable
	case constants.Buyed:
		status = model.StatusArchived
	default:
		status = model.StatusAvailable
	}
	return status, nil
}

func (r *OrderResolver) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*models.Order, error) {
	err := r.orderErrors.CheckCreateOrderInput(input)
	if err != nil {
		return nil, err
	}

	orderCategories := r.orderTypeConverter.ConvertCategoryTypes(input)
	orderCurrency := r.orderTypeConverter.ConvertCurrencyTypes(input)

	// if input.Images == nil {
	// 	return nil, fmt.Errorf("images are required")
	// }

	// compressedImgs := []primitive.Binary{}
	// for _, img := range input.Images {
	// 	compressedImg, err := libs.CompressImage(*img)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	binImg := primitive.Binary{
	// 		Data: compressedImg,
	// 	}

	// 	compressedImgs = append(compressedImgs, binImg)
	// }

	order := &models.Order{
		Id:          primitive.NewObjectID(),
		Title:       *input.Title,
		Description: *input.Description,
		Images:      nil,
		Category:    orderCategories,
		Date:        time.Now(),
		Status:      constants.Available,
		Price:       *input.Price,
		Currency:    orderCurrency,
	}

	order, err = r.orderService.CreateOrder(order)
	if err != nil {
		return nil, fmt.Errorf("server: create order, details: %w", err)
	}

	return order, nil
}

func (r *OrderResolver) Orders(ctx context.Context) ([]*models.Order, error) {
	orders, err := r.orderService.GetOrders()
	if err != nil {
		return nil, fmt.Errorf("server: get orders, details: %w", err)
	}

	var orderPointers []*models.Order
	for _, order := range orders {
		orderPointers = append(orderPointers, &order)
	}

	return orderPointers, nil
}

func (r *OrderResolver) Order(ctx context.Context, id string) (*models.Order, error) {
	order, err := r.orderService.GetOrderById(id)
	if err != nil {
		return nil, fmt.Errorf("server: get order, details: %w", err)
	}

	return order, nil
}

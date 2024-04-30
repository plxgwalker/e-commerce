package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"server.go/graph/generated"
	"server.go/graph/model"
	"server.go/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*models.User, error) {
	return r.UserResolver.CreateUser(ctx, input)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, email string) (*models.User, error) {
	return r.UserResolver.DeleteUser(ctx, email)
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUserInput) (interface{}, error) {
	return r.UserResolver.LoginUser(ctx, input)
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.CreateOrderInput) (*models.Order, error) {
	return r.OrderResolver.CreateOrder(ctx, *input)
}

func (r *mutationResolver) ArchiveOrder(ctx context.Context, id *string) (*models.Order, error) {
	panic(fmt.Errorf("not implemented: ArchiveOrder - archiveOrder"))
}

// ID is the resolver for the id field.
func (r *orderResolver) ID(ctx context.Context, obj *models.Order) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

func (r *orderResolver) Title(ctx context.Context, obj *models.Order) (string, error) {
	panic(fmt.Errorf("not implemented: Title - title"))
}

func (r *orderResolver) Category(ctx context.Context, obj *models.Order) (int, error) {
	panic(fmt.Errorf("not implemented: Category - category"))
}

func (r *orderResolver) Status(ctx context.Context, obj *models.Order) (int, error) {
	panic(fmt.Errorf("not implemented: Status - status"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return r.UserResolver.Users(ctx)
}

func (r *queryResolver) User(ctx context.Context, email string) (*models.User, error) {
	return r.UserResolver.User(ctx, email)
}

func (r *queryResolver) Orders(ctx context.Context) ([]*models.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}

func (r *queryResolver) Order(ctx context.Context, id string) (*models.Order, error) {
	panic(fmt.Errorf("not implemented: Order - order"))
}

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return r.UserResolver.ID(ctx, obj)
}

func (r *userResolver) Orders(ctx context.Context, obj *models.User) ([]*models.Order, error) {
	panic(fmt.Errorf("not implemented: Orders - orders"))
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

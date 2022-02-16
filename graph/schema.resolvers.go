package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"famesensor/go-graphql-jwt/graph/generated"
	"famesensor/go-graphql-jwt/graph/model"
	"famesensor/go-graphql-jwt/models"
	"famesensor/go-graphql-jwt/service"
)

func (r *authOpsResolver) Login(ctx context.Context, obj *model.AuthOps, email string, password string) (interface{}, error) {
	token, err := service.Login(email, password)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (r *authOpsResolver) Register(ctx context.Context, obj *model.AuthOps, input model.RegisterUser) (interface{}, error) {
	if err := service.Register(input); err != nil {
		return nil, err
	}
	return models.RegisterResponse{Status: "success"}, nil
}

func (r *mutationResolver) Auth(ctx context.Context) (*model.AuthOps, error) {
	return &model.AuthOps{}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := service.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user.ToUserGraph(), nil
}

// AuthOps returns generated.AuthOpsResolver implementation.
func (r *Resolver) AuthOps() generated.AuthOpsResolver { return &authOpsResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authOpsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

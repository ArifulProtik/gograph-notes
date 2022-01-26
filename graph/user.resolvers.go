package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/ArifulProtik/gograph-notes/auth"
	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/graph/generated"
	"github.com/ArifulProtik/gograph-notes/graph/model"
	"github.com/ArifulProtik/gograph-notes/services"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.UserRes, error) {
	user, err := services.CreateUser(r.dbclient, input)
	if err != nil {
		msg := err.Error()
		return &model.UserRes{
				Error: &msg,
			},
			nil
	}
	return &model.UserRes{
		User: user,
	}, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, input *model.Login) (*model.LoginRes, error) {
	loginres, err := services.SigninUser(r.dbclient, input)
	if err != nil {
		return &model.LoginRes{}, err
	}
	return loginres, nil
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	tokenData := auth.CtxValue(ctx)
	if tokenData == nil {
		return &ent.User{}, errors.New("Unauthorized")
	}
	user, err := services.GetUserByid(r.dbclient, tokenData.ID)
	if err != nil {
		return &ent.User{}, err
	}
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"math"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ArifulProtik/gograph-notes/auth"
	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/ent/notes"
	"github.com/ArifulProtik/gograph-notes/graph/model"
	"github.com/ArifulProtik/gograph-notes/services"
)

func (r *mutationResolver) CreateNote(ctx context.Context, input *model.NewNote) (*model.ResNotes, error) {
	tokenData := auth.CtxValue(ctx)
	if tokenData == nil {
		return nil, errors.New("Unauthorized")
	}
	user, err := services.GetUserByid(r.dbclient, tokenData.ID)
	if err != nil {
		return nil, graphql.ErrorOnPath(ctx, errors.New("Unauthorized"))
	}
	note, err := services.CreateNote(r.dbclient, input, *user)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &model.ResNotes{
		Note: note,
		User: user,
	}, nil
}

func (r *queryResolver) Mynotes(ctx context.Context, perpage *int, page *int) (*model.MuiltipleNotes, error) {
	newpage := (*page - 1) * *perpage
	tokenData := auth.CtxValue(ctx)
	if tokenData == nil {
		return nil, errors.New("Unauthorized")
	}
	user, err := services.GetUserByid(r.dbclient, tokenData.ID)
	if err != nil {
		return nil, graphql.ErrorOnPath(ctx, errors.New("Unauthorized"))
	}
	data, count, err := services.Mynotes(r.dbclient, *user, *perpage, newpage)
	if err != nil {
		return nil, graphql.ErrorOnPath(ctx, errors.New(err.Error()))
	}
	lastpage := int(math.Ceil(float64(count) / float64(*perpage)))
	return &model.MuiltipleNotes{
		Lastpage: &lastpage,
		Page:     page,
		Notes:    data,
		Perpage:  perpage,
	}, nil
}

func (r *queryResolver) Notes(ctx context.Context, perpage *int, page *int) (*model.MuiltipleNotes, error) {
	newpage := (*page - 1) * *perpage
	count, err := r.dbclient.Notes.Query().Count(context.Background())
	if err != nil {
		return nil, graphql.ErrorOnPath(ctx, errors.New(err.Error()))
	}
	lastpage := int(math.Ceil(float64(count) / float64(*perpage)))
	data, err := r.dbclient.Notes.Query().Order(ent.Desc(notes.FieldCreatedAt)).Limit(*perpage).Offset(newpage).WithAuthor().All(context.Background())
	if err != nil {
		return nil, graphql.ErrorOnPath(ctx, errors.New(err.Error()))

	}
	return &model.MuiltipleNotes{
		Lastpage: &lastpage,
		Page:     page,
		Perpage:  perpage,
		Notes:    data,
	}, nil
}

func (r *queryResolver) Singlenote(ctx context.Context, slug string) (*ent.Notes, error) {
	note, err := services.Singlenote(r.dbclient, slug)
	if err != nil {
		return nil, graphql.ErrorOnPath(ctx, errors.New("Not Found"))
	}
	return note, nil
}

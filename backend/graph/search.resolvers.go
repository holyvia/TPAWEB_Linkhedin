package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/holyvia/gqlgen-todos/graph/generated"
	"github.com/holyvia/gqlgen-todos/graph/model"
	"github.com/holyvia/gqlgen-todos/service"
)

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, keyword string, limit int, offset int) (interface{}, error) {
	var search *model.Search
	search = &model.Search{}

	users, err := service.GetUsersByName(ctx, keyword, limit, offset)
	if err != nil {
		return nil, err
	}
	search.Users = users

	return search, nil
}

// Users is the resolver for the users field.
func (r *searchResolver) Users(ctx context.Context, obj *model.Search) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Posts is the resolver for the posts field.
func (r *searchResolver) Posts(ctx context.Context, obj *model.Search) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// Search returns generated.SearchResolver implementation.
func (r *Resolver) Search() generated.SearchResolver { return &searchResolver{r} }

type searchResolver struct{ *Resolver }

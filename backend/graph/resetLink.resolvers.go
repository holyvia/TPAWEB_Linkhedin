package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/holyvia/gqlgen-todos/graph/model"
	"github.com/holyvia/gqlgen-todos/service"
)

// CreateResetLink is the resolver for the createResetLink field.
func (r *mutationResolver) CreateResetLink(ctx context.Context, userEmail string) (string, error) {
	id := uuid.NewString()

	user, err := service.GetUserByEmail(ctx, userEmail)

	if err != nil {
		return "Error", err
	}
	link := &model.ResetLink{
		ID:     id,
		UserID: user.ID,
		Link:   "http://127.0.0.1:5173/resetpass/" + id,
	}

	if err := r.DB.Create(link).Error; err != nil {
		return "Error", err
	}

	service.SendResetPassword(user.Email, link.Link)
	return "Success", nil
}

// GetResetLink is the resolver for the getResetLink field.
func (r *queryResolver) GetResetLink(ctx context.Context, id string) (*model.ResetLink, error) {
	panic(fmt.Errorf("not implemented"))
}

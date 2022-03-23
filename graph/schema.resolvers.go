package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"dudo/event_service/graph/generated"
	"dudo/event_service/graph/model"
	"fmt"
)

func (r *mutationResolver) CreateEvent(ctx context.Context, input []*model.NewEvent) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

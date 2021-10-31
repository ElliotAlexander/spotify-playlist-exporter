package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/elliotalexander/spotify-playlist-exporter/api/graph/generated"
	"github.com/elliotalexander/spotify-playlist-exporter/lib/auth"
)

func (r *mutationResolver) LoginCallback(ctx context.Context, input *string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Login(ctx context.Context) (string, error) {
	var auth auth.SpotifyAuth
	url := auth.Login()
	return url, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/elliotalexander/spotify-playlist-exporter/api/graph/generated"
	"github.com/elliotalexander/spotify-playlist-exporter/lib/auth"
)

func (r *queryResolver) Login(ctx context.Context) (string, error) {
	var auth auth.SpotifyAuth
	url := auth.Login()
	return url, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

type mutationResolver struct{ *Resolver }

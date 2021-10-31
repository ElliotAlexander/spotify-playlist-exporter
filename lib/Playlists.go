package main

/*
import (
	"context"
	"fmt"
	"log"

	"github.com/zmb3/spotify/v2"
)

func (s *Spotify) retrievePlaylistIDs(opts *spotify.Options) (playlists *spotify.SimplePlaylistPage) {
	user, err := s.client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	playlists, playlistErr := s.client.GetPlaylistsForUserOpt(user.ID, opts)
	if playlistErr != nil {
		log.Fatal(playlistErr)
	}

	return
}

func (s *Spotify) retrievePaginatedPlaylists() (result []spotify.SimplePlaylist) {
	var offset = 0
	var limit = 50

	opts := spotify.RequestOption{}
	opts.Limit = &limit
	opts.Offset = &offset

	for {
		playlists := s.retrievePlaylistIDs(&opts)
		playlistArr := playlists.Playlists
		result = append(result, playlistArr...)

		if len(playlistArr) == limit {
			fmt.Println("Got a new set of playlists")
			fmt.Println("Limit = %i", limit)
			fmt.Println("Offset = %i", offset)
			fmt.Println("Returned = %i", len(playlistArr))
			offset += limit
		} else {
			break
		}

	}
	return
}
*/

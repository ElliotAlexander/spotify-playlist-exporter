package main

import (
    "log"
    "github.com/zmb3/spotify"
    "fmt"
)


func retrievePlaylistIDs(client *spotify.Client, opts *spotify.Options) (playlists *spotify.SimplePlaylistPage) {
    user, err := client.CurrentUser()
    if err != nil {
        log.Fatal(err)
    }

    playlists, playlistErr := client.GetPlaylistsForUserOpt(user.ID, opts)
    if playlistErr != nil {
        log.Fatal(playlistErr)
    }

    return
}


func retrievePaginatedPlaylists(client *spotify.Client) (result []spotify.SimplePlaylist) {
	var offset = 0;
	var limit  = 50;

	opts := spotify.Options{}
	opts.Limit  = &limit
	opts.Offset = &offset

	for {
		playlists := retrievePlaylistIDs(client, &opts)
		playlistArr := playlists.Playlists
		result = append(result, playlistArr...)

		if(len(playlistArr) == limit) {
            fmt.Println("Got a new set of playlists")
            fmt.Println("Limit = %i", limit)
            fmt.Println("Offset = %i", offset)
            fmt.Println("Returned = %i", len(playlistArr))
			offset += limit;
		} else {
			break;
		}

	}
	return
}

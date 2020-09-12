package main

import (
	"log"
	"github.com/zmb3/spotify"
)


func retrievePlaylistIDs(client *spotify.Client) (playlists *spotify.SimplePlaylistPage) {
	user, err := client.CurrentUser()
	if err != nil {
	   log.Fatal(err)
	}
 
	playlists, playlistErr := client.GetPlaylistsForUser(user.ID)
	if playlistErr != nil {
	   log.Fatal(playlistErr)
	}
 
	return
 }
package main

import (
	"fmt"
	"log"
	"net/http"
    "encoding/json"
	"github.com/zmb3/spotify"
)

func handleCallback(client *spotify.Client) {
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You are logged in as:", user.ID)
	simplePlaylistPages := *retrievePlaylistIDs(client)
	playlists := simplePlaylistPages.Playlists
	jsonData, err := json.Marshal(playlists)
	if err != nil {
	   log.Println(err)
	}
	fmt.Println(string(jsonData))
    return
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}

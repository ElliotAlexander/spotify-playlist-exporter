package main

import (
    "encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/zmb3/spotify"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate)
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func main() {
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

    clientID := "3b0e667366d34ff0b6148d3f8e7fb6a3"
    secretKey := "acdb971b0c4645e88c8c0618f7330d09"
    auth.SetAuthInfo(clientID, secretKey)
	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)


    retrievePlaylistIDs(client);
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

func retrievePlaylistIDs(client *spotify.Client) {
   user, err := client.CurrentUser()
   if err != nil {
      log.Fatal(err)
   }
   playlists, err := client.GetPlaylistsForUser(user.ID)
   if err != nil {
      log.Fatal(err)
   }

   jsonData, err := json.Marshal(playlists)
   if err != nil {
      log.Println(err)
   }
   fmt.Println(string(jsonData))
}


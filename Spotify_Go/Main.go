package main

import (
    "encoding/json"
    "os"
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

    clientID := os.Getenv("SPOTIFY_CLIENT_ID")
    secretKey := os.Getenv("SPOTIFY_SECRET_KEY")

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

   fmt.Println(string(dumpToJson(playlists)))
}

func dumpToJson(value interface{}) (jsonData []byte) {
   jsonData, err := json.Marshal(value)
   if err != nil {
      log.Println(err)
   }
   return
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const redirectURI = "http://localhost:8080/callback"

type SpotifyAuth struct {
	state string
	auth  *spotifyauth.Authenticator
}

var (
	ch = make(chan *spotify.Client)
)

func (a *SpotifyAuth) Login() (client *spotify.Client) {
	a.auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
	a.state = "abc123"

	// first start an HTTP server
	http.HandleFunc("/callback", a.completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	url := a.auth.AuthURL(a.state)
	openbrowser(url)

	client = <-ch
	return

}

func (a *SpotifyAuth) completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := a.auth.Token(r.Context(), a.state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != a.state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, a.state)
	}

	// use the token to get an authenticated client
	client := spotify.New(a.auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}

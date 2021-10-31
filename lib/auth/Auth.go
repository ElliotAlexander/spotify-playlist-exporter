package auth

import (
	"log"
	"net/http"

	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

const redirectURI = "http://localhost:8080/callback"

type SpotifyAuth struct {
	state string
	auth  *spotifyauth.Authenticator
}

func (a *SpotifyAuth) Login() (url string) {
	a.auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate))
	a.state = "abc123"

	url = a.auth.AuthURL(a.state)
	return
}

func (a *SpotifyAuth) completeAuth(w http.ResponseWriter, r *http.Request) (*oauth2.Token, error) {
	tok, err := a.auth.Token(r.Context(), a.state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != a.state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, a.state)
	}

	return tok, err
}

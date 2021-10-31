package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zmb3/spotify/v2"
)

type Spotify struct {
	auth   SpotifyAuth
	client *spotify.Client
}

func main() {
	LoadEnv()
	spotify := Spotify{}
	spotify.init()
}

func (s *Spotify) init() {
	s.auth = SpotifyAuth{}
	s.client = s.auth.Login()

	user, err := s.client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You are logged in as:", user.ID)
}

package main

import (
	"github.com/zmb3/spotify/v2"
)

type Spotify struct {
	client *spotify.Client
}

func main() {

}

func (s *Spotify) init() {
	LoadEnv()
}

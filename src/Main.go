package main

import (
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
    go log.Fatal(http.ListenAndServe(":8080", nil))

    clientID := os.Getenv("SPOTIFY_CLIENT_ID")
    secretKey := os.Getenv("SPOTIFY_SECRET_KEY")

    auth.SetAuthInfo(clientID, secretKey)
    url := auth.AuthURL(state)
    openBrowser(url)

    client := <-ch
    playlists := handleCallback(client)

    for _, playlist := range playlists {
        tracks := getSongsForPlaylist(client, playlist.ID)
        fmt.Fprintf(os.Stdout, "\n\nPlaylist Name: %s\n", playlist.Name)
        value := dumpToJson(tracks)
        writeToFile("playlists/", playlist.Name+".json", value)
        for _, track := range tracks {
            artistStr := ""
            for _, artist := range track.Track.Artists {
                artistStr += artist.Name+", "
            }
            fmt.Fprintf(os.Stdout, "\t- %s - %s\n", track.Track.Name, artistStr)
        }
    }

    fmt.Println("\n\n\nDone!")
}

func handleCallback(client *spotify.Client)  (playlistArr []spotify.SimplePlaylist){
    user, err := client.CurrentUser()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("You are logged in as:", user.ID)
    playlistArr = retrievePaginatedPlaylists(client)
    return
}



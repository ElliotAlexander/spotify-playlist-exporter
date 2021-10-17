package main

import (
    "encoding/json"
    "os"
    "fmt"
    "log"
    "net/http"
    "github.com/zmb3/spotify"
    "io/ioutil"
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
	playlists := handleCallback(client)

	jsonData, err := json.Marshal(playlists)
    if err != nil {
        log.Println(err)
    }
    fmt.Println(string(jsonData))

	writeToFile(jsonData)
}

func check(e error) {
    if e != nil {
	    panic(e)
    }
}


func dumpToJson(value interface{}) (jsonData []byte) {
    jsonData, err := json.Marshal(value)
    if err != nil {
        log.Println(err)
    }
    return
}

func writeToFile(data []byte) {
    err := ioutil.WriteFile("/tmp/dat1", data, 0644)
    check(err)
}

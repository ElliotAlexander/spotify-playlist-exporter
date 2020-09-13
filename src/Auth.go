package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/zmb3/spotify"
)

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

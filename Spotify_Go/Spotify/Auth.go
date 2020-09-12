package main

import (
   "github.com/zmb3/spotify"
)

func RetrievePlaylists(client * spotify.Client) (playlists *client.SimplePlaylistsPageForUser) {
   user, err := client.CurrentUser()
   if err != nil {
      log.Fatal(err)
   }
   playlists, err := client.GetPlaylistsForUser(user.ID)
   if err != nil {
      log.Fatal(err)
   }

   return
}



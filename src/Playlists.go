package main

import (
    "log"
    "github.com/zmb3/spotify"
)

func retrievePlaylistIDs(client *spotify.Client, opts *spotify.Options) (playlists *spotify.SimplePlaylistPage) {
    user, err := client.CurrentUser()
    if err != nil {
        log.Fatal(err)
    }

    playlists, playlistErr := client.GetPlaylistsForUserOpt(user.ID, opts)
    if playlistErr != nil {
        log.Fatal(playlistErr)
    }

    return
}


func retrievePaginatedPlaylists(client *spotify.Client) (result []spotify.SimplePlaylist) {
    var offset = 0;
    var limit  = 50;

    opts := spotify.Options{}
    opts.Limit  = &limit
    opts.Offset = &offset

    for {
        playlists := retrievePlaylistIDs(client, &opts)
        playlistArr := playlists.Playlists
        result = append(result, playlistArr...)

        if(len(playlistArr) == limit) {
            offset += limit;
        } else {
            break;
        }

    }
    return
}

func getSongsForPlaylist(client *spotify.Client, playlistId spotify.ID) (result []spotify.PlaylistTrack) {
    var offset = 0
    var limit  = 100
    var fields = "items(track(name,href,artists(name),album(name,href)))"

    opts := spotify.Options{}
    opts.Limit = &limit
    opts.Offset = &offset

    for {
        playlistTrackPage, err := client.GetPlaylistTracksOpt(playlistId, &opts, fields)
        if err != nil {
            panic(err)
        }

        tracks := playlistTrackPage.Tracks
        result = append(result, tracks...)

        if(len(tracks) == limit ) {
            offset += limit
        } else {
            break
        }
    }

    return
}

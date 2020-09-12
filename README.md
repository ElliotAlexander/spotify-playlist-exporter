# Spotify Playlist Exporter

This is a simple Go tool to pull a complete backup of your spotify playlists to JSON. The intention behind this is to have a file-backup of the contents of your spotify playlists (I have several hundred!) in the event that I ever want to move away from spotify, or my account is deleted/hacked/lost/other playlist-data losing event. 

# Setup

To set this up yourself, you'll need to create an App over at [the Spotify Developer Dashboard](https://developer.spotify.com/dashboard/applications), and set your Client ID and Client Secret to two environment variables, `SPOTIFY_CLIENT_ID` and `SPOTIFY_SECRET_KEY` respectively.

Once that's up and running, run the app with:

`go run src/*`



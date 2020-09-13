# Spotify Playlist Exporter

![Lint](https://github.com/ElliotAlexander/Spotify-exporter/workflows/golangci-ling/badge.svg)

This is a simple Go tool to pull a complete backup of your spotify playlists to JSON. The intention behind this is to have a file-backup of the contents of your spotify playlists (I have several hundred!) in the event that I ever want to move away from spotify, or my account is deleted/hacked/lost/other playlist-data losing event. 

## Setup

To set this up yourself, you'll need to create an App over at [the Spotify Developer Dashboard](https://developer.spotify.com/dashboard/applications), and set your Client ID and Client Secret to two environment variables, `SPOTIFY_CLIENT_ID` and `SPOTIFY_SECRET_KEY` respectively.

Once that's up and running, run the app with:

`go run src/*`

Your web browser should open automatically. If not, you'll need to copy-paste the printed URL from the console to your browser. This will log you into Spotify, and start dumping your playlists. Once complete, your complete playlist collection will be dumped to `/playlists` in the repository root.

## Maintenance 

I've no real plans to continue development past simple backups. If a serious spotify competitor emerges, I may look into functionality to reupload a set of downloaded playlist data, but that's likely as far as the roadmap extends.

If you're interested in contributing to, or extending this project in anyway, please do!

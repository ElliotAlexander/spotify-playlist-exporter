module github.com/elliotalexander/spotify-playlist-exporter/api

go 1.17

replace github.com/elliotalexander/spotify_exporter/lib@0.1 => lib@0.1

require (
	github.com/99designs/gqlgen v0.14.0
	github.com/vektah/gqlparser/v2 v2.2.0

)

require (
	github.com/agnivade/levenshtein v1.1.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/golang-lru v0.5.0 // indirect
	github.com/mitchellh/mapstructure v0.0.0-20180203102830-a4e142e9c047 // indirect
)

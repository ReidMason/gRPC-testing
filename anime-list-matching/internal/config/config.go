package config

import "os"

func GetPlexToken() string {
	return os.Getenv("PLEX_TOKEN")
}

func GetPlexURL() string {
	return os.Getenv("PLEX_URL")
}

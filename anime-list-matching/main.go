package main

import (
	"anime-list-matching/internal/anilist"
	"anime-list-matching/internal/animeDb"
	"anime-list-matching/internal/migrations"
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
)

func main() {
	connectionString := "postgres://user:password@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Failed to connect to database", err)
	}

	migrations.ApplyMigrations(db)

	ctx := context.Background()
	queries := animeDb.New(db)

	recurseAnime(16498, make([]anilist.Anime, 0), queries, ctx)

	// animeResult := anilist.GetAnime(113415, queries, ctx)

	// log.Print(animeResult)
}

func recurseAnime(animeId int32, path []anilist.Anime, queries *animeDb.Queries, ctx context.Context) error {
	anime := anilist.GetAnime(animeId, queries, ctx)
	path = append(path, anime)

	sequelId, err := getSequelID(anime.Relations)
	if err != nil {
		printTraversalPath(path)
		return errors.New("No sequel found")
	}

	recurseAnime(sequelId, path, queries, ctx)

	return nil
}

func getSequelID(relations []anilist.Relation) (int32, error) {
	for _, relation := range relations {
		if relation.Relation == anilist.Sequel {
			return relation.ID, nil
		}
	}

	return 0, errors.New("No sequel found")
}

func printTraversalPath(path []anilist.Anime) {
	for i, anime := range path {
		log.Println(strings.Repeat(" ", i), getTraversalPathPrefix(i), anime.Title.Romaji)
	}
}

func getTraversalPathPrefix(index int) string {
	if index == 0 {
		return ""
	}

	return " ∟"
}

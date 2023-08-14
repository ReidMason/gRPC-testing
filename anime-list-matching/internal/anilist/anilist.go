package anilist

import (
	"anime-list-matching/internal/animeDb"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type GraphqlBody[T any] struct {
	Query     string `json:"query"`
	Variables T      `json:"variables"`
}

type GetAnimeVariables struct {
	AnimeId int32 `json:"anime_id"`
}

func GetAnime(animeId int32, queries *animeDb.Queries, ctx context.Context) AnimeResponse {
	res, err := queries.GetCachedAnimeResult(ctx, animeId)
	if err == nil {
		var data AnimeResponse
		err = json.Unmarshal(res, &data)
		if err != nil {
			log.Print(err)
		}

		return data
	}

	log.Println("Not found making request")

	url := "https://graphql.anilist.co/"

	query := `query ($anime_id: Int) {
    Media(id: $anime_id, type: ANIME) {
      id
      format
      episodes
      synonyms
      status
      endDate {
        year
        month
        day
        }
      startDate {
        year
        month
        day
        }
      title {
        english
        romaji
        }
      relations {
        edges {
          relationType
            }
        nodes {
          id
          format
          endDate {
            year
            month
            day
                }
            startDate {
                year
                month
                day
                }
            }
        }
    }
}`

	body := GraphqlBody[GetAnimeVariables]{
		Query: query,
		Variables: GetAnimeVariables{
			AnimeId: animeId,
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Print(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data AnimeResponse
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Print(err)
	}

	_, err = queries.CacheAnimeResult(ctx, animeDb.CacheAnimeResultParams{
		ID:       int32(data.Data.Media.ID),
		Response: jsonData,
	})
	if err != nil {
		log.Println(err)
	}

	return data
}
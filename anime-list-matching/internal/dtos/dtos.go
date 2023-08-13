package dtos

type AnimeResponse struct {
	Data struct {
		Media struct {
			ID       int      `json:"id"`
			Format   string   `json:"format"`
			Episodes int      `json:"episodes"`
			Synonyms []string `json:"synonyms"`
			Status   string   `json:"status"`
			EndDate  struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"endDate"`
			StartDate struct {
				Year  int `json:"year"`
				Month int `json:"month"`
				Day   int `json:"day"`
			} `json:"startDate"`
			Title struct {
				English string `json:"english"`
				Romaji  string `json:"romaji"`
			} `json:"title"`
			Relations struct {
				Edges []struct {
					RelationType string `json:"relationType"`
				} `json:"edges"`
				Nodes []struct {
					ID      int    `json:"id"`
					Format  string `json:"format"`
					EndDate struct {
						Year  any `json:"year"`
						Month any `json:"month"`
						Day   any `json:"day"`
					} `json:"endDate"`
					StartDate struct {
						Year  int `json:"year"`
						Month int `json:"month"`
						Day   int `json:"day"`
					} `json:"startDate"`
				} `json:"nodes"`
			} `json:"relations"`
		} `json:"Media"`
	} `json:"data"`
}

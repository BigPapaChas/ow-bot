package ow

// Profile represents the decoded response from ow-api's /profile endpoint
type Profile struct {
	Name             string        `json:"name"`
	GamesWon         int           `json:"gamesWon"`
	Ratings          []ratings     `json:"ratings"`
	CompetitiveStats gameplayStats `json:"competitiveStats"`
	QuickPlayStats   gameplayStats `json:"quickPlayStats"`
	Level            int           `json:"level"`
	Presitige        int           `json:"prestige"`
	Endorsement      int           `json:"endorsement"`
}

type ratings struct {
	Level int    `json:"level"`
	Role  string `json:"role"`
}

type gameplayStats struct {
	Awards awards `json:"awards"`
	Games  games  `json:"games"`
}

type awards struct {
	Cards        int `json:"cards"`
	Medals       int `json:"medals"`
	MedalsBronze int `json:"medalsBronze"`
	MedalsSilver int `json:"medalsSilver"`
	MedalsGold   int `json:"medalsGold"`
}

type games struct {
	Played int `json:"played"`
	Won    int `json:"won"`
}

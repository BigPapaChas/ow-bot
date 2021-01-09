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

// Heroes represents the decoded response from ow-api's /heroes endpoint
type Heroes struct {
	Name             string        `json:"name"`
	GamesWon         int           `json:"gamesWon"`
	Ratings          []ratings     `json:"ratings"`
	CompetitiveStats detailedStats `json:"competitiveStats"`
	QuickPlayStats   detailedStats `json:"quickPlayStats"`
	Level            int           `json:"level"`
	Presitige        int           `json:"prestige"`
	Endorsement      int           `json:"endorsement"`
}

type detailedStats struct {
	Awards      awards                         `json:"awards"`
	Games       games                          `json:"games"`
	CareerStats map[string]detailedPlayerStats `json:"careerStats"`
	TopHeroes   map[string]playerStats         `json:"topHeroes"`
}

type playerStats struct {
	TimePlayed          string `json:"timePlayed"`
	GamesWon            int    `json:"gamesWon"`
	WinPercentage       int    `json:"winPercentage"`
	WeaponAccuracy      int    `json:"weaponAccuracy"`
	EliminationsPerLife int    `json:"eliminationsPerLife"`
	MultiKillBest       int    `json:"multiKillBest"`
	ObjectiveKills      int    `json:"objectiveKills"`
}

type detailedPlayerStats struct {
	Assists       assists        `json:"assists"`
	Average       average        `json:"average"`
	Best          best           `json:"best"`
	Combat        combat         `json:"combat"`
	HeroSpecific  map[string]int `json:"heroSpecific"`
	Game          game           `json:"game"`
	MatchAwards   matchAwards    `json:"matchAwards"`
	Miscellaneous miscellaneous  `json:"miscellaneous"`
}

type assists struct {
	DefensiveAssists            int `json:"defensiveAssists"`
	DefensiveAssistsAvgPer10Min int `json:"defensiveAssistsAvgPer10Min"`
	DefensiveAssistsMostInGame  int `json:"defensiveAssistsMostInGame"`
	OffensiveAssists            int `json:"offensiveAssists"`
	OffensiveAssistsAvgPer10Min int `json:"offensiveAssistsAvgPer10Min"`
	OffensiveAssistsMostInGame  int `json:"offensiveAssistsMostInGame"`
	ReconAssists                int `json:"reconAssists"`
	ReconAssistsAvgPer10Min     int `json:"reconAssistsAvgPer10Min"`
	ReconAssistsMostInGame      int `json:"reconAssistsMostInGame"`
	HealingDone                 int `json:"healingDone"`
	HealingDoneMostInGame       int `json:"healingDoneMostInGame"`
}

type average struct {
}

type best struct {
}

type combat struct {
}

type game struct {
}

type matchAwards struct {
}

type miscellaneous struct {
}

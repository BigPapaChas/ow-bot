package ow

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var testProfile = Profile{
	Name:     "TestUser",
	GamesWon: 10,
	Ratings: []ratings{
		{Level: 2000, Role: "tank"},
		{Level: 2300, Role: "damage"},
		{Level: 2400, Role: "support"},
	},
	CompetitiveStats: gameplayStats{
		Awards: awards{
			Cards:        100,
			Medals:       12,
			MedalsBronze: 3,
			MedalsSilver: 4,
			MedalsGold:   5,
		},
		Games: games{
			Played: 200,
			Won:    100,
		},
	},
	QuickPlayStats: gameplayStats{
		Awards: awards{
			Cards:        200,
			Medals:       24,
			MedalsBronze: 6,
			MedalsSilver: 8,
			MedalsGold:   10,
		},
		Games: games{
			Played: 400,
			Won:    200,
		},
	},
	Level:       50,
	Presitige:   2,
	Endorsement: 3,
}

func TestProfileApi(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(testProfileHandler))
	defer srv.Close()

	c := NewClient(srv.Client())
	u, _ := url.Parse(srv.URL)
	c.BaseURL = u

	result, err := c.GetProfile()
	if err != nil {
		t.Errorf("Unexpected error on request: %s", err)
	}
	if !reflect.DeepEqual(*result, testProfile) {
		t.Error("Profiles didn't match", testProfile, "!=", result)
	}
}

func testProfileHandler(w http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal(testProfile)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(data)
	if err != nil {
		panic(err)
	}
}

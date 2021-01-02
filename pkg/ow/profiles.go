package ow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// RegisteredProfile struct
type RegisteredProfile struct {
	BattlenetID string
	Platform    string
}

// TODO: Swap to
// Structs for the `profile` endpoint of the ow-api
type player struct {
	Ratings []rating `json:"ratings"`
}

type rating struct {
	Level int64  `json:"level"`
	Role  string `json:"role"`
}

// PrintRatings prints the SR rating of each role
func (r *RegisteredProfile) PrintRatings() {
	if ratings, err := r.getRatings(); err == nil {
		fmt.Println(r.BattlenetID, ratings)
	} else {
		fmt.Println("Failed to retrieve ratings for", r.BattlenetID)
	}
}

// TODO:
func (r *RegisteredProfile) getRatings() ([]rating, error) {

	// ow-api requires the battlenet tag to be of the form ProfileName-1234
	urlSafeID := strings.Replace(r.BattlenetID, "#", "-", 1)
	url := fmt.Sprintf("https://ow-api.com/v1/stats/%s/us/%s/profile", r.Platform, urlSafeID)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var p = new(player)
	err = json.Unmarshal(body, &p)
	if err != nil {
		return nil, err
	}

	return p.Ratings, nil
}

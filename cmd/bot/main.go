package main

import (
	"context"
	"log"
	"os"

	"github.com/owbot/internal/db"
)

func main() {
	projectID := os.Getenv("PROJECT_ID")
	profilesCollection := os.Getenv("PROFILES_COLLECTION")
	statsCollection := os.Getenv("STATS_COLLECTION")
	collectionPrefix := os.Getenv("COLLECTION_PREFIX")

	log.Printf(projectID)
	log.Printf(profilesCollection)
	log.Printf(statsCollection)
	log.Printf(collectionPrefix)

	dbClient, err := db.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatal(err)
	}
	dbClient.InitCollections(collectionPrefix+profilesCollection, collectionPrefix+statsCollection)

	p := db.Profile{
		UserID:     "1234",
		Username:   "BigPapaChas#8534",
		BattleTags: []string{"BigPapaChas#11352", "Blah#12345"},
	}

	err = dbClient.UpdateProfile(p, "1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully created profile:\n %+v\n", p)

	if err != nil {
		log.Fatal(err)
	}

	//id := "1234"

	// p := db.Profile{
	// 	UserID:     id,
	// 	Username:   "BigPapaChas#8534",
	// 	BattleTags: []string{"BigPapaChas#11352", "Blah#12345"},
	// }

	err = dbClient.RemoveBattleTag("1234", "ChrisTag#1234")
	if err != nil {
		log.Fatal(err)
	}

	// tags, err := dbClient.GetBattleTags(id)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("User with id %s has the following battletags: %v", id, tags)
	// ranks := []db.Rankings{
	// 	{
	// 		Damage:  db.Rank{Rating: 1000},
	// 		Support: db.Rank{Rating: 2000},
	// 	},
	// 	{
	// 		Support: db.Rank{Rating: 3000},
	// 	},
	// 	{
	// 		Tank:    db.Rank{Rating: 1000},
	// 		Damage:  db.Rank{Rating: 1000},
	// 		Support: db.Rank{Rating: 2000},
	// 	},
	// 	{
	// 		Damage: db.Rank{Rating: 1000},
	// 	},
	// }
	// for index, t := range tags {
	// 	rank := ranks[index%len(ranks)]
	// 	err = dbClient.UpdateCompRankings(t, rank)
	// 	if err != nil {
	// 		log.Print(err)
	// 	} else {
	// 		log.Printf("Updated %s's competitive rankings to %+v", t, rank)
	// 	}
	// }

	// for _, t := range tags {
	// 	r, err := dbClient.GetCompRankings(t)
	// 	if err != nil {
	// 		log.Print(err)
	// 	} else {
	// 		log.Printf("%s's competitive rankings are %+v", t, r)
	// 	}
	// }
}

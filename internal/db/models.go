package db

type Profile struct {
	UserID     string   `firestore:"userID"`
	Username   string   `firestore:"username"`
	BattleTags []string `firestore:"battleTags"`
}

type Stats struct {
	Competitive Rankings `firestore:"competitive"`
	Casual      Rankings `firestore:"casual"`
}

type Rankings struct {
	Tank    Rank `firestore:"tank"`
	Damage  Rank `firestore:"damage"`
	Support Rank `firestore:"support"`
}

type Rank struct {
	Rating int `firestore:"rating"`
}

package db

import (
	"context"

	"cloud.google.com/go/firestore"
)

type Client struct {
	fsClient   *firestore.Client
	profileRef *firestore.CollectionRef
	statsRef   *firestore.CollectionRef
	context    context.Context
}

func NewClient(ctx context.Context, projectID string) (*Client, error) {
	fsClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	dbClient := &Client{
		fsClient: fsClient,
		context:  ctx,
	}
	return dbClient, nil
}

func (c *Client) Close() error {
	return c.fsClient.Close()
}

func (c *Client) InitCollections(profileCollection, statsCollection string) {
	c.profileRef = c.fsClient.Collection(profileCollection)
	c.statsRef = c.fsClient.Collection(statsCollection)
}

func (c *Client) GetBattleTags(userID string) ([]string, error) {
	p := &Profile{}
	err := c.loadProfile(userID, p)
	if err != nil {
		return nil, err
	}
	return p.BattleTags, nil
}

func (c *Client) GetCompRankings(battleTag string) (*Rankings, error) {
	s := &Stats{}
	err := c.loadStats(battleTag, s)
	if err != nil {
		return nil, err
	}
	return &s.Competitive, nil
}

func (c *Client) UpdateCompRankings(battleTag string, rankings Rankings) error {
	var paths []firestore.FieldPath
	if rankings.Damage.Rating != 0 {
		paths = append(paths, firestore.FieldPath{"competitive", "damage", "rating"})
	}
	if rankings.Tank.Rating != 0 {
		paths = append(paths, firestore.FieldPath{"competitive", "tank", "rating"})
	}
	if rankings.Support.Rating != 0 {
		paths = append(paths, firestore.FieldPath{"competitive", "support", "rating"})
	}

	opts := firestore.Merge(paths...)
	stats := Stats{Competitive: rankings}
	_, err := c.statsRef.Doc(battleTag).Set(c.context, stats, opts)
	return err
}

func (c *Client) UpdateProfile(p Profile, userID string) error {
	_, err := c.profileRef.Doc(userID).Set(c.context, p)
	return err
}

func (c *Client) loadProfile(userID string, p *Profile) error {
	return c.loadDocument(c.profileRef, userID, p)
}

func (c *Client) loadStats(battleTag string, s *Stats) error {
	return c.loadDocument(c.statsRef, battleTag, s)
}

func (c *Client) loadDocument(colRef *firestore.CollectionRef, key string, data interface{}) error {
	docsnap, err := colRef.Doc(key).Get(c.context)
	if err != nil {
		return err
	}
	return docsnap.DataTo(&data)
}

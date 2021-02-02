package db

import (
	"context"

	"cloud.google.com/go/firestore"
)

const (
	profilesCollectionName = "profiles"
	statsCollectionName    = "stats"
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

func (c *Client) InitCollections(collectionPrefix string) {
	c.profileRef = c.fsClient.Collection(collectionPrefix + profilesCollectionName)
	c.statsRef = c.fsClient.Collection(collectionPrefix + statsCollectionName)
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

func (c *Client) GetCompDamage(battleTag string) (*Rank, error) {
	s := &Stats{}
	err := c.loadStats(battleTag, s)
	if err != nil {
		return nil, err
	}
	return &s.Competitive.Damage, nil
}

func (c *Client) GetCompSupport(battleTag string) (*Rank, error) {
	s := &Stats{}
	err := c.loadStats(battleTag, s)
	if err != nil {
		return nil, err
	}
	return &s.Competitive.Damage, nil
}

func (c *Client) GetCompTank(battleTag string) (*Rank, error) {
	s := &Stats{}
	err := c.loadStats(battleTag, s)
	if err != nil {
		return nil, err
	}
	return &s.Competitive.Tank, nil
}

func (c *Client) UpdateProfile(p Profile, userID string) error {
	_, err := c.profileRef.Doc(userID).Set(c.context, p)
	return err
}

func (c *Client) AddBattleTag(userID string, battleTag string) error {
	p := &Profile{}
	err := c.loadProfile(userID, p)
	if err != nil {
		return err
	}
	for _, t := range p.BattleTags {
		if t == battleTag {
			return nil
		}
	}
	p.BattleTags = append(p.BattleTags, battleTag)
	_, err = c.profileRef.Doc(userID).Set(c.context, p)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) RemoveBattleTag(userID string, battleTag string) error {
	p := &Profile{}
	err := c.loadProfile(userID, p)
	if err != nil {
		return err
	}
	for i, t := range p.BattleTags {
		if t == battleTag {
			p.BattleTags = append(p.BattleTags[:i], p.BattleTags[i+1:]...)
			return c.UpdateProfile(*p, userID)
		}
	}
	return nil
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

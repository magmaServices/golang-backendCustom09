package model

import (
	"encoding/json"

	"github.com/gocraft/dbr"
)

const (
	tableHeroes = "heroes"
)

// Hero defines a single record in heroes table
type Hero struct {
	ID        int    `db:"hero_id"`
	HeroName  string `db:"hero_name"`
	PlayerID  int    `db:"player_id"`
	HeroStats string `db:"hero_stats"`
}

func InsertHero(tx *dbr.Tx, hero *Hero) error {
	r, err := tx.
		InsertInto(tableHeroes).
		Columns(
			// "hero_id",
			"hero_name",
			"player_id",
			"hero_stats",
		).
		Record(hero).
		Exec()
	if err != nil {
		return err
	}
	i, err := r.LastInsertId()
	if err != nil {
		return err
	}
	hero.ID = int(i)
	return nil
}

// FindHeroesByPlayerID returns a list of heroes associated with specified player
func (q *Queries) FindHeroesByPlayerID(sess *dbr.Session, playerID int) (hs []Hero, err error) {
	_, err = sess.
		Select(
			"hero_id",
			"hero_name",
			"player_id",
		).
		From(tableHeroes).
		Where("player_id = ?", playerID).
		Load(&hs)

	return hs, err
}

// FindHeroByName returns a hero with specified name
func (q *Queries) FindHeroByName(sess *dbr.Session, heroName string) (h Hero, err error) {
	err = sess.
		Select(
			"hero_id",
			"hero_name",
			"player_id",
		).
		From(tableHeroes).
		Where("hero_name = ?", heroName).
		LoadOne(&h)
	return h, err
}

// FindHeroStats returns stats of hero of specified ID
func (q *Queries) FindHeroStats(sess *dbr.Session, heroID int) (pr HeroStats, err error) {
	var payload []byte
	err = sess.
		Select("hero_stats").
		From(tableHeroes).
		Where("hero_id = ?", heroID).
		LoadOne(&payload)
	if err != nil {
		return pr, err
	}

	err = json.Unmarshal(payload, &pr)
	return pr, err
}

// UpdateHeroStats changes stats of hero of specified ID
func (q *Queries) UpdateHeroStats(tx *dbr.Tx, heroID int, pr *HeroStats) error {
	by, err := json.Marshal(pr)
	if err != nil {
		return err
	}
	_, err = tx.
		Update(tableHeroes).
		Set("hero_stats", string(by)).
		Where("hero_id = ?", heroID).
		Exec()
	return err
}

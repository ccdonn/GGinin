package entity

import (
	"time"
)

type Review struct {
	Id      string    `db:"Rvw_ID"`
	Uid     string    `db:"Poster_ID"`
	Content string    `db:"Content"`
	Create  time.Time `db:"Create_Time"`
	Display time.Time `db:"Display_Time"`
	Url1    string    `db:"ImageUrl_1"`
	Url2    string    `db:"ImageUrl_2"`
	Url3    string    `db:"ImageUrl_3"`
	Url4    string    `db:"ImageUrl_4"`
	Url5    string    `db:"ImageUrl_5"`
	Spot    string    `db:"Spot_ID"`
	Privacy int       `db:"Privacy"`
	Score   float32   `db:"Score"`
	Alert   int       `db:"Alert"`
	Status  int       `db:"Status_ID"`
}

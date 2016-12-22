package entity

import ()

type Article struct {
	Id      int `db:"Article_ID"`
	Created int
	Title   string
	Content string
}

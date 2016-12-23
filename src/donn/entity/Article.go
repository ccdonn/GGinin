package entity

import ()

type Article struct {
	Id      int64 `db:"Article_ID"`
	Created int64
	Title   string
	Content string
}

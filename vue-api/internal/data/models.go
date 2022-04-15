package data

import (
	"database/sql"
	"time"
)

const DB_TIME_OUT = 3 * time.Second

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		User:   User{},
		Token:  Token{},
		Book:   Book{},
		Author: Author{},
	}
}

type Models struct {
	User   User
	Token  Token
	Book   Book
	Author Author
	Genre  Genre
}

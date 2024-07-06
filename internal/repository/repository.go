package repository

import "github.com/jmoiron/sqlx"

type UserRepository struct {
	conn *sqlx.DB
}

func NewUser(conn *sqlx.DB) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

type WorkRepository struct {
	conn *sqlx.DB
}

func NewWork(conn *sqlx.DB) *WorkRepository {
	return &WorkRepository{
		conn: conn,
	}
}

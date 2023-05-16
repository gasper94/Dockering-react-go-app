package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface{
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

// docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
func NewPostgresStore()(*PostgresStore, error){
	// connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	url := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=%v",
    "postgres",
    "gobank",
    "0.0.0.0",
    "5432",
    "postgres",
	"disable")

	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s * PostgresStore) CreateAccount(*Account) error {
	return nil
}

func (s * PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s * PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s * PostgresStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
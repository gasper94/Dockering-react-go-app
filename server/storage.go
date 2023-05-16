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
	GetAccounts() ([]*Account, error)
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

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s * PostgresStore) CreateAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		created_at timestamp
	)`	

	_, err := s.db.Exec(query)
	return err
}

func (s * PostgresStore) CreateAccount(acc *Account) error {
	query := `insert into account 
	(first_name, last_name, created_at)
	values ($1, $2, $3)`

	_, err := s.db.Query(
		query, 
		acc.FirstName, 
		acc.LastName,
		acc.CreatedAt,
	)

	if err != nil {
		return err
	}


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

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from account")

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next(){
		account := new(Account)
		err := rows.Scan(
			&account.ID, 
			&account.FirstName, 
			&account.LastName,
			&account.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		
		accounts = append(accounts, account)
	}

	return accounts, nil
}

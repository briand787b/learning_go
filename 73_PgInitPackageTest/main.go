package main

import (
	"database/sql"
	"log"
	"github.com/briand787b/pgInit"
	"fmt"
	"os"
	"time"
)

type User struct {
	ID		string
	Email 		string
	HashedPassword 	string
	Username	string
	CreatedAt 	time.Time
	ModifiedAt 	time.Time
}

type UserStore interface {
	Find(int) (*User, error)
	FindByEmail(string) (*User, error)
	FindByUsername(string) (*User, error)
	Save(*User) error
}

type DBUserStore struct {
	db *sql.DB
}

var globalUserStore UserStore

func (store DBUserStore) Save(user *User) error {
	rows, err := store.db.Query(
		`INSERT INTO usr
		(username, email, hashed_password, created_at, modified_at)
		VALUES
		($1, $2, $3, $4, $5)
		RETURNING id;`,
		user.Username,
		user.Email,
		user.HashedPassword,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (store DBUserStore) Find(id int) (*User, error) {
	var usr User
	rows, err := store.db.Query(`SELECT id, username, email, hashed_password, created_at, modified_at FROM usr WHERE id = $1`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&usr.ID, &usr.Username, &usr.Email, &usr.HashedPassword, &usr.CreatedAt, &usr.ModifiedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return &usr, nil
}

func (store DBUserStore) FindByEmail(email string) (*User, error) {
	if email == "" {
		return nil, nil
	}

	var usr User
	rows, err := store.db.Query(`SELECT id, username, email, hashed_password, created_at, modified_at FROM usr WHERE email = $1`, email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&usr.ID, &usr.Username, &usr.Email, &usr.HashedPassword, &usr.CreatedAt, &usr.ModifiedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return &usr, nil
}

func (store DBUserStore) FindByUsername(username string) (*User, error) {
	if username == "" {
		return nil, nil
	}

	var usr User
	rows, err := store.db.Query(`SELECT id, username, email, hashed_password, created_at, modified_at FROM usr WHERE username = $1`, username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&usr.ID, &usr.Username, &usr.Email, &usr.HashedPassword, &usr.CreatedAt, &usr.ModifiedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return &usr, nil
}

func NewDBUserStore(db *sql.DB) (*DBUserStore) {
	return &DBUserStore{db: db}

}

func main() {
	db, err := pgInit.ConnectDefault("CardsForChan")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	globalUserStore = NewDBUserStore(db)
	usr, err := globalUserStore.Find(3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(usr)

	user := User{
		Username: "brianD",
		Email:	"brianD@gmail.com",
		HashedPassword: "13241324234324324",
	}
	err = globalUserStore.Save(&user)
	fmt.Println(err)
	fmt.Println(user.ID)
}

package restdb

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Hostname = "localhost"
	Port     = 5432
	Username = "mtsouk"
	Password = "pass"
	Database = "restapi"
)

type User struct {
	ID        int
	Username  string
	Password  string
	LastLogin int64
	Admin     int
	Active    int
}

func ConnectPostgres() (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", Username, Password, Hostname, Port, Database)

	return pgxpool.New(context.Background(), dbUrl)
}

func DeleteUser(pool *pgxpool.Pool, ID int) error {
	_, err := FindUserID(pool, ID)
	if err != nil {
		return err
	}

	tx, err := pool.Begin(context.Background())
	if err != nil {
		return err
	}
	_, err = tx.Exec(context.Background(), "DELETE FROM users WHERE ID = $1", ID)
	if err != nil {
		tx_err := tx.Rollback(context.Background())
		if tx_err != nil {
			return tx_err
		}
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func ListAllUsers(pool *pgxpool.Pool) ([]User, error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM users\n")
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err := rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			return []User{}, err
		}
		temp := User{ID: c1, Username: c2, Password: c3, LastLogin: c4, Admin: c5, Active: c6}
		all = append(all, temp)
	}
	log.Println("All:", all)

	return all, nil
}

func FindUserID(pool *pgxpool.Pool, ID int) (User, error) {
	rows, err := pool.Query(context.Background(), "SELECT * from users WHERE ID = $1\n", ID)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	u := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err := rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			return User{}, err
		}
		u = User{ID: c1, Username: c2, Password: c3, LastLogin: c4, Admin: c5, Active: c6}
		log.Println("Found user:", u)
	}

	if u.ID == 0 {
		return User{}, fmt.Errorf("user %d does not exists", ID)
	}

	return u, nil
}

func IsUserValid(pool *pgxpool.Pool, u User) (bool, error) {
	rows, err := pool.Query(context.Background(), "SELECT username, password FROM users WHERE username = $1", u.Username)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	var username, password string

	for rows.Next() {
		err := rows.Scan(&username, &password)
		if err != nil {
			return false, err
		}
	}

	if u.Username == username && u.Password == password {
		return true, nil
	}

	return false, nil
}

type UserAlreadyExists struct {
	Username string
}

func (error UserAlreadyExists) Error() string {
	return fmt.Sprintf("User %s already exists", error.Username)
}

func InsertUser(pool *pgxpool.Pool, u User) error {
	ok, err := IsUserValid(pool, u)
	if err != nil {
		return err
	}
	if ok {
		return UserAlreadyExists{Username: u.Username}
	}

	tx, err := pool.Begin(context.Background())
	if err != nil {
		return err
	}
	_, err = tx.Exec(context.Background(), "INSERT INTO users (username, password, lastlogin, admin, active) VALUES ($1, $2, $3, $4, $5)", u.Username, u.Password, u.LastLogin, u.Admin, u.Active)
	if err != nil {
		tx_err := tx.Rollback(context.Background())
		if tx_err != nil {
			return tx_err
		}
		return err
	}
	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func FindUserUsername(pool *pgxpool.Pool, username string) (User, error) {
	rows, err := pool.Query(context.Background(), "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	u := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err := rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			return User{}, err
		}
		u = User{ID: c1, Username: c2, Password: c3, LastLogin: c4, Admin: c5, Active: c6}
		log.Println("Found user:", u)
	}

	return u, nil
}

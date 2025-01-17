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

	_, err = pool.Exec(context.Background(), "DELETE FROM users WHERE ID = $1", ID)
	if err != nil {
		return err
	}

	log.Println("Deleted user with ID:", ID)

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

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/testaquatic/Mastering_GO/ch10/restapi/restdb"
)

type User struct {
	ID        int
	Username  string
	Password  string
	LastLogin int64
	Admin     int
	Active    int
}

var (
	Hostname = "localhost"
	Port     = 5432
	Username = "mtsouk"
	Password = "pass"
	Database = "restapi"
)

func main() {
	pool, err := restdb.ConnectPostgres()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Close()

	err = pool.Ping(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	t := restdb.User{}
	fmt.Println(t)
	rows, err := pool.Query(context.Background(), "SELECT username from users")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(username)
	}

	log.Println("Populating PostgreSQL")
	user := restdb.User{
		ID:        0,
		Username:  "mtsouk",
		Password:  "admin",
		LastLogin: time.Now().Unix(),
		Admin:     1,
		Active:    1,
	}
	err = restdb.InsertUser(pool, user)
	if err == nil {
		log.Println("User inserted successfully.")
	} else {
		log.Println("Insert failed:", err)
	}

	mtsoukUser, err := restdb.FindUserUsername(pool, user.Username)
	if err != nil {
		log.Println("User not found:", err)
	} else {
		log.Println("User found:", mtsoukUser)
	}

	err = restdb.DeleteUser(pool, mtsoukUser.ID)
	if err != nil {
		log.Println("Delete failed:", err)
	} else {
		log.Println("User deleted successfully.")
	}

	mtsoukUser, err = restdb.FindUserUsername(pool, user.Username)
	if err == nil {
		log.Println("User Deleted.")
	} else {
		log.Println("User not Deleted", err)
	}
	mtsoukUser, err = restdb.FindUserUsername(pool, user.Username)
	if err == nil {
		log.Println("User Deleted.")
	} else {
		log.Println("User not Deleted", err)
	}

	err = restdb.DeleteUser(pool, mtsoukUser.ID)
	if err != nil {
		log.Println("User Deleted", err)
	} else {
		log.Println("User not Deleted")
	}
}

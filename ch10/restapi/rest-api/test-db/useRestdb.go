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
	Username = "postgres"
	Password = "pass"
	Database = "restapi"
)

func main() {
	pool := restdb.ConnectPostgres()
	fmt.Println(pool)
	defer pool.Close()

	err := pool.Ping(context.Background())
	if err != nil {
		fmt.Println("Ping:", err)
		return
	}

	t := restdb.User{}
	fmt.Println(t)
	rows, err := pool.Query(context.Background(), "SELECT username FROM users")
	if err != nil {
		fmt.Println("Query:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		err = rows.Scan(&username)
		if err != nil {
			fmt.Println("Scan:", err)
			return
		}
		fmt.Println(username)
	}

	log.Println("Popualating PostgreSQL")
	user := restdb.User{
		ID:        0,
		Username:  "mtsouk",
		Password:  "admin",
		LastLogin: time.Now().Unix(),
		Admin:     1,
		Active:    1,
	}
	if pool.InsertUser(user) {
		log.Println("User inserted successfully.")
	} else {
		log.Println("Insert failed!")
	}

	mtsoukUser := pool.FindUserUsername(user.Username)
	fmt.Println("mtsouk: ", mtsoukUser)

	if pool.DeleteUser(mtsoukUser.ID) {
		log.Println("User deleted.")
	} else {
		log.Println("User not Deleted.")
	}

	mtsoukUser = pool.FindUserUsername(user.Username)
	fmt.Println("mtsouk: ", mtsoukUser)

	if pool.DeleteUser(mtsoukUser.ID) {
		log.Println("User deleted.")
	} else {
		log.Println("User not Deleted.")
	}

	if pool.DeleteUser(mtsoukUser.ID) {
		log.Println("User deleted.")
	} else {
		log.Println("User not Deleted.")
	}
}

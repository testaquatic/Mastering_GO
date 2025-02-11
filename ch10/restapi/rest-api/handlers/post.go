package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/testaquatic/Mastering_GO/ch10/restapi/restdb"
)

func AddHandler(pgpool *restdb.PgPool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("AddHandler Serving:", r.URL.Path, "from", r.Host)
		d, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("AddHandler:", err)
			return
		}

		if len(d) == 0 {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("AddHandler: No input!")
			return
		}

		var users = []restdb.User{}
		err = json.Unmarshal(d, &users)
		if err != nil {
			log.Println("AddHandler:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println(users)
		if !pgpool.IsUserAdmin(users[0]) {
			log.Println("Issued by non-admin user:", users[0].Username)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		result := pgpool.InsertUser(users[1])
		if !result {
			rw.WriteHeader(http.StatusBadRequest)
		}
	}
}

func LoginHandler(pgpool *restdb.PgPool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("LoginHandler Serving:", r.URL.Path, "from", r.Host)

		d, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("LoginHandler:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(d) == 0 {
			log.Println("LoginHandler: No input!")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var user = restdb.User{}
		err = json.Unmarshal(d, &user)
		if err != nil {
			log.Println("LoginHandler:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("Input user:", user)

		if !pgpool.IsUserValid(user) {
			log.Println("LoginHandler: User", user.Username, "not valid!")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		t := pgpool.FindUserUsername(user.Username)
		log.Println("Logged in:", t)

		t.LastLogin = time.Now().Unix()
		t.Active = 1
		if pgpool.UpdateUser(t) {
			log.Println("User updated:", t)
			rw.WriteHeader(http.StatusOK)
		} else {
			log.Println("Update failed:", t)
			rw.WriteHeader(http.StatusBadRequest)
		}
	}
}

func LogoutHandler(pgpool *restdb.PgPool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("LogoutHandler Serving:", r.URL.Path, "from", r.Host)

		d, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("LogoutHandler:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if len(d) == 0 {
			log.Println("LogoutHandler: No input!")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var user = restdb.User{}
		err = json.Unmarshal(d, &user)
		if err != nil {
			log.Println("LogoutHandler:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if !pgpool.IsUserValid(user) {
			log.Println("LogoutHandler: User", user.Username, "not valid!")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("Input user:", user)

		t := pgpool.FindUserUsername(user.Username)
		log.Println("Logged out:", t.Username)
		t.Active = 0
		if pgpool.UpdateUser(t) {
			log.Println("User updated:", t)
			rw.WriteHeader(http.StatusOK)
		} else {
			log.Println("Update failed:", t)
			rw.WriteHeader(http.StatusBadRequest)
		}
	}
}

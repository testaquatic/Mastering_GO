package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/testaquatic/Mastering_GO/ch10/restapi/restdb"
)

func UpdateHandler(pgpool *restdb.PgPool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("UpdateHandler Serving:", r.URL.Path, "from", r.Host)
		d, err := io.ReadAll(r.Body)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("UpdateHandler:", err)
			return
		}

		if len(d) == 0 {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("UpdateHandler: No input!")
			return
		}

		var users = []restdb.User{}
		err = json.Unmarshal(d, &users)
		if err != nil {
			log.Println("UpdateHandler:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if !pgpool.IsUserAdmin(users[0]) {
			log.Println("Command issued by non-admin user:", users[0].Username)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println(users)
		t := pgpool.FindUserUsername(users[1].Username)
		t.Username = users[1].Username
		t.Password = users[1].Password
		t.Admin = users[1].Admin

		if !pgpool.UpdateUser(t) {
			log.Println("Update failed:", t)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("Update successful:", t)
		rw.WriteHeader(http.StatusOK)
	}
}

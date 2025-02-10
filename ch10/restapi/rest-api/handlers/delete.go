package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/testaquatic/Mastering_GO/ch10/restapi/restdb"
)

func DeleteHandler(pgpool *restdb.PgPool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("DeleteHandler Serving:", r.URL.Path, "from", r.Host)

		// 삭제할 유저ID를 얻는다.
		id, ok := mux.Vars(r)["id"]
		if !ok {
			log.Println("DeleteHandler: ID value not set!")
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		var user = restdb.User{}
		err := user.FromJSON(r.Body)
		defer r.Body.Close()
		if err != nil {
			log.Println("DeleteHandler:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		if !pgpool.IsUserAdmin(user) {
			log.Println("DeleteHandler: User", user.Username, "is not admin!")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		intID, err := strconv.Atoi(id)
		if err != nil {
			log.Println("DeleteHandler:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		t := pgpool.FindUserID(intID)
		if t.Username != "" {
			log.Println("About to delete:", t)
			deleted := pgpool.DeleteUser(intID)
			if deleted {
				log.Println("User deleted:", id)
				rw.WriteHeader(http.StatusOK)
				return
			}
		} else {
			log.Println("DeleteHandler: User ID not found:", id)
			rw.WriteHeader(http.StatusNotFound)
			return
		}
	}
}

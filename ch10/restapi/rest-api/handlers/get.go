package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/testaquatic/Mastering_GO/ch10/restapi/restdb"
)

func GetAllHandler(pgpool *restdb.PgPool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("GetAllHandler Serving:", r.URL.Path, "from", r.URL.Host)
		d, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("GetAllHandler Error:", err)
			return
		}
		if len(d) == 0 {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("GetAllHandler Error: empty No input!")
			return
		}

		var user = restdb.User{}
		err = json.Unmarshal(d, &user)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("GetAllHandler Error:", err)
			return
		}

		// 관리자의 권한 확인
		if !pgpool.IsUserAdmin(user) {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("GetAllHandler Error:", user, "is not admin!")
			return
		}

		err = SliceToJSON(pgpool.ListAllUsers(), rw)
		if err != nil {
			log.Println("GetAllHandler Error:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

func GetIDHandler(pgpool *restdb.PgPool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("GetIDHandler Serving:", r.URL.Path, "from", r.URL.Host)

		username, ok := mux.Vars(r)["username"]
		if !ok {
			rw.WriteHeader(http.StatusNotFound)
			log.Println("GetIDHandler Error: ID value not set!")
			return
		}

		d, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("GetIDHandler Error:", err)
			return
		}
		if len(d) == 0 {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("GetIDHandler Error: No input!")
			return
		}

		var user = restdb.User{}
		err = json.Unmarshal(d, &user)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("GetIDHandler Error:", err)
			return
		}

		log.Println("Input user:", user)
		if !pgpool.IsUserAdmin(user) {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("GetIDHandler Error:", "User", user.Username, "not an admin!")
			return
		}

		t := pgpool.FindUserUsername(username)
		if t.ID != 0 {
			err = t.ToJSON(rw)
			if err != nil {
				log.Println("GetIDHandler Error:", err)
				rw.WriteHeader(http.StatusBadRequest)
			}
		} else {
			rw.WriteHeader(http.StatusNotFound)
			log.Println("GetIDHandler Error:", "User", username, "not found!")
		}
	}
}

func GetUserDataHandler(pgpool *restdb.PgPool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("GetUserDataHandler Serving:", r.URL.Path, "from", r.URL.Host)
		id, ok := mux.Vars(r)["id"]
		if !ok {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("GetUserDataHandler Error: ID value not set!")
			return
		}

		intID, err := strconv.Atoi(id)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("GetUserDataHandler Error:", err)
			return
		}

		t := pgpool.FindUserID(intID)
		if t.ID != 0 {
			err = t.ToJSON(rw)
			if err != nil {
				log.Println("GetUserDataHandler Error:", err)
				rw.WriteHeader(http.StatusBadRequest)
			}
			return
		}

		rw.WriteHeader(http.StatusNotFound)
		log.Println("GetUserDataHandler Error:", "User not found:", id)
	}
}

func LoggedUserHandler(pgpool *restdb.PgPool) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("LoggedUserHandler Serving:", r.URL.Path, "from", r.URL.Host)
		var user = restdb.User{}
		err := user.FromJSON(r.Body)
		defer r.Body.Close()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			log.Println("LoggedUserHandler Error:", err)
			return
		}

		if !pgpool.IsUserValid(user) {
			log.Println("LoggedUserHandler Error:", "User", user.Username, "not valid!")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = SliceToJSON(pgpool.ReturnLoggedUsers(), rw)
		if err != nil {
			log.Println("LoggedUserHandler Error:", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

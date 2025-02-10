package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/testaquatic/Mastering_GO/ch10/restapi/rest-api/handlers"
	"github.com/testaquatic/Mastering_GO/ch10/restapi/restdb"
)

// 서버 정보
var rMux = mux.NewRouter()
var PORT = ":1234"

func main() {
	flag.Parse()
	if flag.NArg() >= 1 {
		PORT = ":" + flag.Arg(0)
	}

	pgpool, err := restdb.ConnectPostgres()
	if err != nil {
		log.Println("Error connecting to database:", err)
		return
	}
	defer pgpool.Close()

	s := http.Server{
		Addr:         PORT,
		Handler:      rMux,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	rMux.NotFoundHandler = http.HandlerFunc(handlers.DefaultHandler)

	notAllowed := handlers.NotAllowedHandler{}
	rMux.MethodNotAllowedHandler = notAllowed

	rMux.HandleFunc("/time", handlers.TimeHandler)

	// GET 메서드를 사용하는 함수를 등록한다.
	getMux := rMux.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/getall", handlers.GetAllHandler(pgpool))
	getMux.HandleFunc("/getid/{username}", handlers.GetIDHandler(pgpool))
	getMux.HandleFunc("/logged", handlers.LoggedUserHandler(pgpool))
	getMux.HandleFunc("/username/{id:[0-9]+}", handlers.GetUserDataHandler(pgpool))

	// PUT 메서드를 사용하는 함수를 등록한다.
	putMux := rMux.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/update", handlers.UpdateHandler(pgpool))

	// POST 메서드를 사용하는 함수를 등록한다.
	postMux := rMux.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/login", handlers.LoginHandler(pgpool))
	postMux.HandleFunc("/add", handlers.AddHandler(pgpool))
	postMux.HandleFunc("/logout", handlers.LogoutHandler(pgpool))

	// DELETE 메서드를 사용하는 함수를 등록한다.
	deleteMux := rMux.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/username/{id:[0-9]+}", handlers.DeleteHandler(pgpool))

	// 서버 실행
	go func() {
		log.Println("Listening to", PORT)
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			return
		}
	}()

	// 시그널 처리
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	sig := <-sigs
	log.Println("Qutting after signal:", sig)
	_ = s.Shutdown(context.TODO())
	time.Sleep(5 * time.Second)
}

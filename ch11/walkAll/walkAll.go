package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {}

type notAllowedHandler struct{}

func (h notAllowedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler(w, r)
}

func main() {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(handler)
	nowAllowed := notAllowedHandler{}
	r.MethodNotAllowedHandler = nowAllowed

	// GET
	getMux := r.Methods(http.MethodGet).Subrouter()
	getMux.HandleFunc("/time", handler)
	getMux.HandleFunc("/getall", handler)
	getMux.HandleFunc("/getid", handler)
	getMux.HandleFunc("/logged", handler)
	getMux.HandleFunc("/username/{id:[0-9]+}", handler)

	// PUT
	putMux := r.Methods(http.MethodPut).Subrouter()
	putMux.HandleFunc("/update", handler)

	// POST
	postMux := r.Methods(http.MethodPost).Subrouter()
	postMux.HandleFunc("/add", handler)
	postMux.HandleFunc("/login", handler)
	postMux.HandleFunc("/logout", handler)

	// DELETE
	deleteMux := r.Methods(http.MethodDelete).Subrouter()
	deleteMux.HandleFunc("/username/{id:[0-9]+}", handler)

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		qT, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(qT, ","))
		}
		qRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(qRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ", "))
		}
		fmt.Println()

		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	http.Handle("/", r)
}

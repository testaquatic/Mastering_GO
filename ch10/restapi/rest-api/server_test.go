package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/testaquatic/Mastering_GO/ch10/restapi/rest-api/handlers"
	"github.com/testaquatic/Mastering_GO/ch10/restapi/restdb"
)

func TestTimeHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.TimeHandler)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, "/time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.MethodNotAllowedHandler)

	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestLogin(t *testing.T) {
	UserPass := []byte(`{"Username":"admin","Password":"admin"}`)

	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	pgpool, err := restdb.ConnectPostgres()
	if err != nil {
		t.Fatal(err)
	}
	defer pgpool.Close()

	handler := http.HandlerFunc(handlers.LoginHandler(pgpool))
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAdd(t *testing.T) {
	now := int(time.Now().Unix())
	username := "test_" + strconv.Itoa(now)
	users := `[{"Username":"admin","Password":"admin"},{"Username":"` + username + `","Password":"myPass"}]`

	UserPass := []byte(users)
	req, err := http.NewRequest(http.MethodPost, "/add", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	pgpool, err := restdb.ConnectPostgres()
	if err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(handlers.AddHandler(pgpool))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestGetUserDataHandler(t *testing.T) {
	UserPass := []byte(`{"Username":"admin", "Password":"admin"}`)
	req, err := http.NewRequest(http.MethodGet, "/username/1", bytes.NewBuffer(UserPass))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	pgpool, err := restdb.ConnectPostgres()
	if err != nil {
		t.Fatal(err)
	}
	defer pgpool.Close()
	handler := http.HandlerFunc(handlers.GetUserDataHandler(pgpool))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		return
	}

	expected := `{"ID":1,"Username":"admin","Password":"admin","LastLogin":0,"Admin":1,"Active":1}`
	serveResponse := rr.Body.String()

	result := strings.Split(serveResponse, "LastLogin")
	serveResponse = result[0] + `LastLogin":0,"Admin":1,"Active":1}`

	if serveResponse != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

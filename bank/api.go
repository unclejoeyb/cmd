package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string
}
type APIServer struct {
	listenaddr string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func newApiServer(listenaddr string) *APIServer {
	return &APIServer{
		listenaddr: listenaddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))
	
	log.Println("http server running on port: ", s.listenaddr)

	http.ListenAndServe(s.listenaddr, router)
}

func (s *APIServer) handleAccount (w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s APIServer) handleGetAccount (w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	// account := newAccount("John", "Doe")
	log.Println(id)
	return writeJSON(w, http.StatusOK, &Account{})
}

func (s APIServer) handleCreateAccount (w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s APIServer) handleDeleteAccount (w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s APIServer) handleTransfer (w http.ResponseWriter, r *http.Request) error {
	return nil
}


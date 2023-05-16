package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
// The main function to start the API server
func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	// Starting the server
	log.Println("JSON API SERVER RUNNING ON PORT: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

// Function to handle requests to /account
func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET"{
		return s.handleGetAccount(w,r)
	}

	if r.Method == "POST"{
		return s.handleCreateAccount(w,r)
	}

	
	if r.Method == "DELETE"{
		return s.handleDeleteAccount(w,r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

/* ACCOUNT */
// Function to handle requests to /account
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("Ulises", "Martinez")

	return WriteJSON(w, http.StatusOK, account);
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// HELPER FUNCTIONS
// A helper function to write JSON responses to the client
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// The type of function that handles API requests
type apiFunc func(http.ResponseWriter, *http.Request) error

// The error response format
type ApiError struct {
	Error string `json:"error"`
}

// A function to convert an API function into a function that can be handled by the HTTP router
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if err := f(w,r); err != nil{
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

// The API server structure
type APIServer struct {
	listenAddr string
}

// Function to create a new API server
func NewApiServer(listenAddr string) *APIServer {
	return &APIServer {
		listenAddr: listenAddr,
	}
}
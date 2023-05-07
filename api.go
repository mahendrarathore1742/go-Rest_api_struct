package main

import (
	"encoding/json"
	"net/http"
)

type apieFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPhander(f apieFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if err := f(w, r); err != nil {

			if e, ok := err.(apierror); ok {
				writeJSON(w, e.Status, e)
				return
			}

			writeJSON(w, http.StatusInternalServerError, apierror{Err: "Interna Error", Status: http.StatusInternalServerError})

		}
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request) error {

	if r.Method != http.MethodGet {
		return apierror{Err: "invalid ", Status: http.StatusMethodNotAllowed}
	}

	user := User{}

	user.valid = true

	if !user.valid {
		return apierror{Err: "User is not valid", Status: http.StatusForbidden}
	}

	return writeJSON(w, http.StatusOK, User{})

}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)

}

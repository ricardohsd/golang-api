package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type jsonError struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

const jsonHeader string = "application/json;charset=UTF-8"

func JsonHeader(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsonHeader)

		inner.ServeHTTP(w, r)
	})
}

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func UsersCreate(w http.ResponseWriter, r *http.Request) {
	var user User

	body, err := ioutil.ReadAll(io.Reader(r.Body))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	processJson := func(body []byte, user User) (int, interface{}) {
		if err := json.Unmarshal(body, &user); err != nil {
			json := jsonError{Code: http.StatusBadRequest, Text: "payload has wrong format"}

			return http.StatusUnprocessableEntity, json
		} else {
			u := RepoCreate(user)

			return http.StatusCreated, u
		}
	}

	header, response := processJson(body, user)

	w.WriteHeader(header)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

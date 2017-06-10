package main

type User struct {
	Name string `json:"name,omitempty"`
}

type Users []User

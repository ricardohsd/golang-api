package main

import (
	"testing"
)

func TestRepoCreate(t *testing.T) {
	users = []User{}

	newUser := User{Name: "Mary"}

	RepoCreate(newUser)

	if !containsUser(users, "Mary") {
		t.Errorf("Users list '%v' should contain '%v", users, newUser)
	}
}

func containsUser(list []User, name string) bool {
	for _, u := range list {
		if u.Name == name {
			return true
		}
	}
	return false
}

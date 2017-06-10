package main

var users Users

func init() {
	users = []User{}
}

func RepoCreate(u User) User {
	users = append(users, u)
	return u
}

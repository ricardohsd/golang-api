package main

var users Users

func init() {
	users = []User{}
}

func RepoCreate(u User) User {
	if !isAlreadyCreated(u) {
		users = append(users, u)
	}
	return u
}

func isAlreadyCreated(newUser User) bool {
	for _, u := range users {
		if u.Name == newUser.Name {
			return true
		}
	}
	return false
}

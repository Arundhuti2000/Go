package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	User := User{}
	User.Name = name
	User.Type = membershipType
	if membershipType == "premium" {
		User.MessageCharLimit = 1000
	} else {
		User.MessageCharLimit = 100
	}
	return User
}

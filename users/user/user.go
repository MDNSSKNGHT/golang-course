package user

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func New(firstName, lastName, birthDate string) User {
	return User{firstName, lastName, birthDate, time.Now()}
}

func (user User) PrintDetails() {
	fmt.Printf("Hello, %s %s\n", user.firstName, user.lastName)
	fmt.Printf("Your birthdate is %s\n", user.birthDate)
	fmt.Printf("Your account was created at %s\n", user.createdAt.Format(time.DateTime))
}

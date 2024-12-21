package main

import "fmt"

type User struct {
	id   uint64
	name string
}

func (u *User) UpdateName(newName string) {
	u.name = newName

}

func struct1() {
	user := User{
		id:   1,
		name: "Luiz",
	}

	user.UpdateName("Carlos")
	fmt.Println(user.id, user.name)
}

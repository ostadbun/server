package main

import (
	"fmt"
	"ostadbun/entity"
)

func main() {

	user := entity.User{Email: "alirezak"}

	myuser := user.Hash()

	fmt.Printf("%s\n", myuser.Email)
}

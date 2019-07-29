package main

import (
	"encoding/json"
	"fmt"
	"github.com/wawandco/fako"
)

type User struct {
	Id       int
	Name     string `fako:"full_name"`
	Username string `fako:"user_name"`
	Email    string `fako:"email_address"` //Notice the fako:"email_address" tag
	Phone    string `fako:"phone"`
	Password string `fako:"simple_password"`
	Address  string `fako:"street_address"`
}

func main() {
	for i := 1; i <= 100000; i++ {
		var user User
		fako.Fill(&user)
		user.Id = i
		js, _ := json.Marshal(user)
		fmt.Println(string(js))
	}
}

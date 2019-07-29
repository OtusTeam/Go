package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type User struct {
	Id       int
	Name     string `fako:"full_name"`
	Username string `fako:"user_name"`
	Email    string `fako:"email_address"`
	Phone    string `fako:"phone"`
	Password string `fako:"simple_password"`
	Address  string `fako:"street_address"`
}

func getComDomains(filename string) map[string]uint32 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("%v", err)
	}

	lines := strings.Split(string(fileContents), "\n")

	users := make([]User, 0)

	for _, line := range lines {
		user := &User{}
		json.Unmarshal([]byte(line), user)

		users = append(users, *user)
	}

	comDomains := make(map[string]uint32)

	for _, user := range users {
		matched, err := regexp.Match("\\.com", []byte(user.Email))
		if err != nil {
			log.Fatalf("%v", err)
		}

		if matched {
			num := comDomains[strings.SplitN(user.Email, "@", 2)[1]]
			num++
			comDomains[strings.SplitN(user.Email, "@", 2)[1]] = num
		}
	}

	return comDomains
}

func main() {
	comDomains := getComDomains("data.dat")
	log.Printf("%v", comDomains)
}

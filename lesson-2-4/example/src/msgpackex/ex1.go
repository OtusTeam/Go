package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
)

type Person struct {
	Name        string
	Surname     string
	Age         uint8
	ChildrenAge map[string]uint8
}

func main() {
	p := Person{
		Name:  "Ivan",
		Surname: "Remen",
		Age:   27,
	}
	p.ChildrenAge = make(map[string]uint8)
	p.ChildrenAge["Alex"] = 5
	p.ChildrenAge["Maria"] = 2

	marshaled, _ := msgpack.Marshal(&p)

	fmt.Printf("Length of marshaled: %v IMPL: %v\n", len(marshaled), string(marshaled))

	p2 := &Person{}
	msgpack.Unmarshal(marshaled, p2)
	fmt.Printf("Unmarshled: %v\n", p2)
}

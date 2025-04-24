package main

import (
	"fmt"
	"log"

	"github.com/buskari/golang-oop/structs"
)

func main() {
	p, err := structs.NewPerson("A", "Buscariolli", 36)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
	p.HaveABirthday()
	fmt.Println(p)
}

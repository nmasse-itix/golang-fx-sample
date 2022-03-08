package main

import (
	"os"

	lib "github.com/nmasse-itix/golang-fx"
)

func main() {
	gm := lib.NewCat("Gros Minet")
	isidore := lib.NewCat("Isidore")
	c1, err := lib.NewChild(2)
	if err != nil {
		panic(err)
	}
	c2, err := lib.NewChild(5)
	if err != nil {
		panic(err)
	}
	john := lib.NewAdult("John", []lib.Child{c1, c2})
	jane := lib.NewAdult("Jane", []lib.Child{c1, c2})
	house := lib.NewHouse("New-York", []lib.Adult{john, jane}, []lib.Cat{gm, isidore})

	house.Present(os.Stdout)
}

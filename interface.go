package main

import "fmt"

type Movement interface {
	move()
}

//interface
type Animal interface {
	speak()
}
type Dog struct{}

// method implement animal
func (d Dog) speak() {
	fmt.Println("goooooooooooooooo..........")
}

func (d Dog) move() {
	fmt.Println("runnnnnnnnnn...........")
}

// embed inteface
type NextAnimal interface {
	Movement
	Animal
}

func main() {
	fmt.Println("==================================")

	dog := Dog{}

	var m Movement = dog
	m.move()

	var a Animal = dog
	a.speak()

	// var na = NextAnimal = dog
	// na.speak()
	// na.move()
}

package main

import (
	"fmt"
	"strings"
)

const (
	_       = 0
	cow int = iota
	bird
	snake
)

const (
	_       = 0
	eat int = iota
	move
	speak
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func NewCow() *Cow {
	return &Cow{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}
}

func (c *Cow) Eat() {
	fmt.Println(c.food)
}

func (c *Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c *Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func NewBird() *Bird {
	return &Bird{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}
}

func (b *Bird) Eat() {
	fmt.Println(b.food)
}

func (b *Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b *Bird) Speak() {
	fmt.Println(b.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func NewSnake() *Snake {
	return &Snake{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}
}

func (s *Snake) Eat() {
	fmt.Println(s.food)
}

func (s *Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s *Snake) Speak() {
	fmt.Println(s.noise)
}

func main() {

	animals := make(map[string]int)

	fmt.Println("=======\nPress Ctrl-c to stop")
	for {
		var action int
		fmt.Println("\nChoose an action:\n1) newanimal\n2) query")
		fmt.Print("> ")
		fmt.Scan(&action)

		if action == 1 {
			fmt.Println("New animal creation")
			saveAnimal(animals)
		} else if action == 2 {
			fmt.Println("Animal query")
			animalName := askName()
			if !isAnimalExist(animalName, animals) {
				fmt.Printf("I don't know the animal with the name: %s", animalName)
				continue
			}
			queryAnimal(animalName, animals)
		}

	}

}

func askName() string {
	var animalName string
	fmt.Println("\nPlease, type an animal Name:")
	fmt.Print(" > ")
	fmt.Scan(&animalName)

	return strings.ToLower(animalName)
}

func askType() int {
	var animalType int
	fmt.Println("\nPlease, choose an animal Type:\n1) cow\n2) bird\n3) snake")
	fmt.Print(" > ")
	fmt.Scan(&animalType)

	return animalType
}

func askQuery() int {
	var query int
	fmt.Println("\nWhat information would you like to know:\n1) eat\n2) move\n3) speak")
	fmt.Print(" > ")
	fmt.Scan(&query)

	return query
}

func saveAnimal(animals map[string]int) {

	animalName := askName()
	animalType := askType()

	switch animalType {
	case cow:
		animals[animalName] = cow
	case bird:
		animals[animalName] = bird
	case snake:
		animals[animalName] = snake
	}

	fmt.Println("\nCreated it!")
}

func createAnimal(animalType int) Animal {

	switch animalType {
	case cow:
		return NewCow()
	case bird:
		return NewBird()
	case snake:
		return NewSnake()
	}
	return nil
}

func isAnimalExist(name string, animals map[string]int) bool {
	_, ok := animals[name]
	return ok
}

func queryAnimal(name string, animals map[string]int) {
	query := askQuery()

	var animalType int
	if !isAnimalExist(name, animals) {
		fmt.Printf("I don't know the animal with the name: %s", name)
	} else {
		animalType = animals[name]
	}

	animal := createAnimal(animalType)

	switch query {
	case move:
		fmt.Printf("Your animal %s can ", strings.Title(name))
		animal.Move()
	case eat:
		fmt.Printf("Your animal %s prefers to eat ", strings.Title(name))
		animal.Eat()
	case speak:
		fmt.Printf("Your animal %s sound is ", strings.Title(name))
		animal.Speak()
	default:
		fmt.Println("I can't do such an action. Try: eat, move, speak")
	}

}

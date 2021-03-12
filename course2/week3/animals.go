package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	animals := make(map[string]Animal)

	animals["cow"] = Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	animals["bird"] = Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	animals["snake"] = Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	fmt.Println("=======\nPress Ctrl-C to stop")
	for {
		var animal string
		fmt.Println("\nChoose an animal: cow, bird, snake")
		fmt.Print("> ")
		fmt.Scan(&animal)

		var info string
		fmt.Println("info: eat, move, speak")
		fmt.Printf("%s > ", animal)
		fmt.Scan(&info)

		if chosen, ok := animals[animal]; !ok {
			fmt.Println("I don't know such an animal. Try again.")
			continue
		} else {
			switch info {
			case "eat":
				chosen.Eat()
			case "move":
				chosen.Move()
			case "speak":
				chosen.Speak()
			default:
				fmt.Println("I don't know what to do")
			}
		}
	}

}

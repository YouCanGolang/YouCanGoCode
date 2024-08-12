package main

import (
	"fmt"
)

type Animal interface {
	Speak()
}

type Dog struct {
	Voice string
}

func (d Dog) Speak() {
	fmt.Println(d.Voice)
}

type Cat struct {
	Voice string
}

func (c Cat) Speak() {
	fmt.Println(c.Voice)
}

func main() {
	var a interface{}
	a = 8
	fmt.Println(a)

	dog := Dog{
		Voice: "Wang...",
	}
	dog.Speak()

	var animalDog Animal = Dog{
		Voice: "Wang...",
	}
	animalDog.Speak()

	var animalCat Animal = Cat{
		Voice: "Miao...",
	}
	animalCat.Speak()

	var animalDog1 Animal = Dog{
		Voice: "Wang...",
	}
	animalSpeak(animalDog1)

	var animalCat1 Animal = Cat{
		Voice: "Miao...",
	}
	animalSpeak(animalCat1)
}

func animalSpeak(animal Animal) {
	animal.Speak()
}

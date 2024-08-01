package main

import "fmt"

type Animal interface {
	Says()
	LikesWater() bool
}

type Dog struct{}

func (d Dog) Says() {
	fmt.Println("Woof!")
}

func (d Dog) LikesWater() bool {
	return true
}

type Cat struct{}

func (c Cat) Says() {
	fmt.Println("Meow!")
}

func (c Cat) LikesWater() bool {
	return false
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct{}

func (df DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (cf CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	dogFactory := DogFactory{}
	catFactory := CatFactory{}

	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	cat.Says()

	fmt.Println("A dog likes water:", dog.LikesWater())
	fmt.Println("A cat likes water:", cat.LikesWater())
}

package main

import (
	"fmt"

	"github.com/gucio321/terminalmenu/constructor"
)

func main() {
	constructor.Create("Test Menu").
		MainPage("Main page").
		Item("Hello World", func() {
			fmt.Println("Hello World")
		}).
		Subpage("Test Page").Item("print test", func() { fmt.Println("test") }).Back().
		Exit().
		Run()
}

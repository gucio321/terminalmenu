package main

import (
	"fmt"

	"github.com/gucio321/terminalmenu/constructor"
)

func main() {
	constructor.Create("Test Menu").Page("Test page").Item("Hello World", func() { fmt.Println("Hello World") }).Back().Run()
}

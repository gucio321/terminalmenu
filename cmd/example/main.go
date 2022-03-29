package main

import (
	"fmt"

	terminalmenu "github.com/gucio321/terminalmenu/pkg"
)

func main() {
	<-terminalmenu.Create("Test Menu", true).
		MainPage("Main page").
		Item("Hello World", func() {
			fmt.Println("Hello World")
		}).
		Subpage("Test Page").Item("print test", func() { fmt.Println("test") }).Back().
		Exit().
		Run()
}

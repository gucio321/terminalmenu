package main

import (
	"fmt"

	terminalmenu "github.com/gucio321/terminalmenu/pkg"
	"github.com/gucio321/terminalmenu/pkg/menuutils"
)

func main() {
	<-terminalmenu.Create("Test Menu", true).
		MainPage("Main page").
		Item("Hello World", func() {
			fmt.Println("Hello World")
			menuutils.PromptEnter("Press ENTER to continue")
		}).
		Subpage("Test Page").Item("print test", func() {
		fmt.Println("test")
		menuutils.PromptEnter("Press ENTER to continue")
	}).Back().
		Exit().
		Run()
}

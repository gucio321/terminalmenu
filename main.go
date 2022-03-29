package main

import "fmt"

func main() {
	menu := Create("Test Menu").Page("Test page").Back()
	fmt.Println(menu)
}

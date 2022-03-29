package main

import "fmt"

func main() {
	menu := Create("Test Menu")
	menu.Page("main", Page("Test page"))
	fmt.Println(menu)
}

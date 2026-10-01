package main

import "fmt"

func main() {
	//var keys map[<key>string]<value>string -> Empty map

	//colors := make(map[<key>string]<value>string)

	/* colors := map[<key>string]<value>string{ -> Whit data map
		"<key>":   "<value>",
	} */

	/* colors := map[string]string{
		"red":   "#A62B2B",
		"green":"#60A62B",
	} */

	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["red"] = "#A62B2B"
	colors["green"] = "#60A62B"

	//delete(colors, "red")

	

	printMap(colors)
}

func printMap(c map[string]string){
	for key, value := range c{
		fmt.Println(key, " -> ", value)
	}
}

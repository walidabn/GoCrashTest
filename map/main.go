package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	//var colors map[string] string
	//colors := make(map[string]string)
	colors["white"] = "#ffffff"
	//	delete(colors, "white")
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c { // color, hex == key,value for range creation
		fmt.Printf("Hex code for %v is %v \n", color, hex)
	}
}

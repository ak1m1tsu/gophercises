package maps

import "fmt"

func GetColors() map[string]string {
	return map[string]string{
		"red":   "#ff00000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
}

func PrintColors(colors map[string]string) {
	for color, hex := range colors {
		fmt.Printf("Hex code for %v is %v\n", color, hex)
	}
}

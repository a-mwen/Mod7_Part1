package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What is your favorite color?") // Asking the first question
	in := bufio.NewReader(os.Stdin)             // Input
	color, _ := in.ReadString('\n')             // Reading it
	color = strings.TrimSpace(color)
	fmt.Println("Your favorite color is", color+"!")

	fmt.Println("What is your favorite food?") // Asking the second question
	food, _ := in.ReadString('\n')             // Reading it
	food = strings.TrimSpace(food)
	fmt.Println("Your favorite food is", food+"!")
}

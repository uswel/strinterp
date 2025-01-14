package main

import (
	"fmt"

	"github.com/uswel/strinterp"
)

func main() {
	variables := map[string]interface{}{
		"name":  "Alice",
		"age":   30,
		"price": 12.345,
		"city":  "New York",
	}

	message, err := strinterp.String("My name is ${name} and I am ${age} years old. The price is ${price:.2f}. $$ is the escape sequence for $. I live in ${city:upper}", variables)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(message)
}

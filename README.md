# strinterp

`strinterp` is a Go library for performing string interpolation, allowing you to embed and format variables directly within strings. It offers a simple and readable syntax inspired by shell scripting and other template languages.

## Features

-   **Simple variable expansion:** `${variable}`
-   **Formatting:** `${variable:format}`
    -   **Integer Formats:**
        -   `d`: Decimal (e.g., `${age:d}`)
        -   `x`: Lowercase hexadecimal (e.g., `${id:x}`)
        -   `X`: Uppercase hexadecimal (e.g., `${id:X}`)
        -   `o`: Octal (e.g., `${permissions:o}`)
        -   `b`: Binary (e.g., `${flags:b}`)
    -   **Float Formats:**
        -   `f`: Fixed-point (e.g., `${price:f}`)
        -   `.<precision>f`: Fixed-point with specified precision (e.g., `${price:.2f}`)
        -   `.<precision>`: Fixed-point with specified precision (shorthand for `.precisionf`) (e.g., `${score:.3}`)
        -   `e`: Lowercase scientific notation (e.g., `${value:e}`)
        -   `E`: Uppercase scientific notation (e.g., `${value:E}`)
        -   `g`: General format (uses `e` or `f` depending on magnitude) (e.g., `${number:g}`)
        -   `G`: General format (uses `E` or `f` depending on magnitude) (e.g., `${number:G}`)
    -   **String Formats:**
        -   `upper`: Uppercase (e.g., `${city:upper}`)
        -   `lower`: Lowercase (e.g., `${name:lower}`)
        -   `title`: Title case (e.g., `${title:title}`)
-   **Escaping:** `$$` to escape the dollar sign
-   **Error handling:**
    -   Returns an error for unclosed variable expansions.
    -   Returns an error for undefined variables.
    -   Returns an error for invalid or unsupported format specifiers.

## Installation

```bash
go get github.com/uswel/strinterp]([invalid URL removed])
````

## Usage

```go
package main

import (
	"fmt"

	"[github.com/uswel/strinterp [invalid URL removed]"
)

func main() {
	variables := map[string]interface{}{
		"name":  "Alice",
		"age":   30,
		"price": 12.345,
		"city":  "New York",
		"id":    255,
		"score": 0.9876,
	}

	message, err := strinterp.String("My name is ${name} and I am ${age} years old. The price is ${price:.2f}. $$ is the escape sequence for $. I live in ${city:upper}. ID: ${id:X}, Score: ${score:.3}", variables)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(message)
	// Output: My name is Alice and I am 30 years old. The price is 12.35. $ is the escape sequence for $. I live in NEW YORK. ID: FF, Score: 0.988
}
```

## API Documentation

### `func String(input string, variables map[string]interface{}) (string, error)`

Expands variables within a string using the syntax `${var}` or `${var:format}`.

**Parameters:**

  - `input`: The string containing variable expansions.
  - `variables`: A map of variable names to their values.

**Returns:**

  - The expanded string with variables replaced by their values.
  - An error if there is an unclosed variable expansion, an undefined variable, or an invalid format specifier.

### `func parseVar(expression string) (name string, format string)`

(Internal function) Extracts the variable name and optional format from the expression.

### `func formatValue(value interface{}, format string) (string, error)`

(Internal function) Formats the variable value based on the optional format.

### `func formatInt(value int64, format string) (string, error)`

(Internal function) Formats an integer value.

### `func formatFloat(value float64, format string) (string, error)`

(Internal function) Formats a float value.

### `func formatString(value string, format string) (string, error)`

(Internal function) Formats a string value.

## Error Handling

The `String` function returns an error in the following cases:

  - **Unclosed variable expansion:** If a variable expansion is not properly closed with a `}`, e.g., `${name`.
  - **Undefined variable:** If a referenced variable is not found in the `variables` map.
  - **Invalid format specifier:** If an unsupported or invalid format specifier is used for a given variable type.
  - **Invalid precision:** if the precision provided in the format specifier is not a valid integer.

## Contributing

Contributions are welcome\! Please see the [CONTRIBUTING.md](/CONTRIBUTING.md) file for details.

## License

This project is licensed under the MIT License - see the [LICENSE](/LICENSE) file for details.

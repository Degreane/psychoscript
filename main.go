package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := `
Set x to 2     # automatically sets x as Int32 with value 2
Set y to 1.3  # automatically sets y as float64 with value 1.3
Set m to x+y - 3
Set n to 1.9 as Integer
Set nn to "this is a # line "
Set nm as String
Set mm as Boolean
Def z as Integer # automatically Defines z as Int32 and initializes it to null value
If z is Null then Print(x) else Print( (x+y)*3).  # Checks if z is null then prints the value of x else prints the value of (x+y)*3 , note here that since x is Int32 and y is float64 then x is temporary changed into float64 then added to y
`

	// Remove comments and trim spaces
	lines := strings.Split(input, "\n")
	var cleanedLines []string
	for _, line := range lines {
		if idx := strings.Index(line, "#"); idx != -1 {
			line = line[:idx]
		}
		line = strings.TrimSpace(line)
		if line != "" {
			cleanedLines = append(cleanedLines, line)
		}
	}
	cleanedInput := strings.Join(cleanedLines, "\n")

	_, err := Parse("", []byte(cleanedInput))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
		os.Exit(1)
	}

	// Print final variable states
	fmt.Println("\nFinal variable states:")
	for name, variable := range variables {
		switch variable.Type {
		case TypeInt:
			fmt.Printf("%s: %d (Int32)\n", name, variable.Int)
		case TypeFloat:
			fmt.Printf("%s: %f (Float64)\n", name, variable.Float)
		case TypeString:
			fmt.Printf("%s: \"%s\" (String)\n", name, variable.String)
		case TypeBoolean:
			fmt.Printf("%s: %t (Boolean)\n", name, variable.Boolean)
		case TypeNull:
			fmt.Printf("%s: Null\n", name)
		}
	}
}

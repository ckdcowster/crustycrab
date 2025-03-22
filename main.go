/*

Hands-on exercise #49 - map[string][]string
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores their favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.

key 				value
`bond_james` 		`shaken, not stirred`, `martinis`, `fast cars`
`moneypenny_jenny` 	`intelligence`, `literature`, `computer science`
`no_dr` 			`cats`, `ice cream`, `sunsets`

*/

package main

import "fmt"

func main() {
	my_map := (make(map[string][]string))

	my_map[`bond_james`] = []string{`shaken`, `not stirred`, `martinis`, `fast cars`}
	my_map[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	my_map[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	for k, v := range my_map {
		fmt.Println("key is", k)
		for i, s := range v {
			fmt.Printf("idx = %v \tvalue = %v\n", i, s)
		}
	}
}

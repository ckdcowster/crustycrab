/*

Hands-on exercise #50 - add a record
Using the code from the previous example, add a record to your map. Now print the map out
using the “range” loop
key 			value
`fleming_ian` 	`steaks`, `cigars`, `espionage`

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

	fmt.Println("-------------------------------")
	my_map[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	for k, v := range my_map {
		fmt.Println("key is", k)
		for i, s := range v {
			fmt.Printf("idx = %v \tvalue = %v\n", i, s)
		}
	}
}

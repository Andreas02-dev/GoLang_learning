package main

import (
	"fmt"
)

func main() {
	stringSlice := []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`}
	stringSlice2 := []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`}
	stringSlice3 := []string{`no_dr`, `Being evil`, `Ice cream`, `Sunsets`}
	myMap := map[string][]string{`bond_james`: stringSlice, `moneypenny_miss`: stringSlice2, `no_dr`: stringSlice3}
	myMap["Andreas"] = []string{"Pizza", "me", "myself"}
	for k, v := range myMap {
		fmt.Println("Dit is de record voor: ", k)
		for i, vv := range v {
			fmt.Printf("Index: %d\t Value: %v\n", i, vv)
		}
	}
}

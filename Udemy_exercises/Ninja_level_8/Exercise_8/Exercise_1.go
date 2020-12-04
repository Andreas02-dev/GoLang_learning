package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	// First Marshall the user slice:
	bs, err := json.Marshal(users)
	// Check for errors:
	if err != nil {
		fmt.Println(err)
	}
	// Convert bs to a string and print it:
	fmt.Println(string(bs))
}

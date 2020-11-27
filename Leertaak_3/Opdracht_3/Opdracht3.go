package main

import (
	"fmt"
)

func main() {
	isBoterAanwezig := true
	isKaasAanwezig := true
	isMelkAanwezig := false
	isEiAanwezig := false

	// Print of er boter én kaas aanwezig is (true):
	fmt.Println(isBoterAanwezig && isKaasAanwezig)
	// Print of er melk én ei aanwezig is (false):
	fmt.Println(isMelkAanwezig && isEiAanwezig)
	// Print of er boter én melk aanwezig is (false):
	fmt.Println(isBoterAanwezig && isMelkAanwezig)
	// Print of er boter of melk aanwezig is (true):
	fmt.Println(isBoterAanwezig || isMelkAanwezig)
	// Print of er boter of kaas aanwezig is (true):
	fmt.Println(isBoterAanwezig || isKaasAanwezig)
	// Print of er melk of ei aanwezig is (false):
	fmt.Println(isMelkAanwezig || isEiAanwezig)
	// Print of er geen melk en wel ei aanwezig is (false):
	fmt.Println(!(isMelkAanwezig) && isEiAanwezig)
	// Print of er zowel melk als ei én geen kaas aanwezig is (false):
	// Gedachteproces hierbij: melk, ei, kaas zijn allemaal NIET aanwezig.
	fmt.Println(!(isMelkAanwezig) && !(isEiAanwezig) && !(isKaasAanwezig))
	// Wanneer melk en ei wel aanwezig moeten zijn en kaas niet aanwezig (false):
	fmt.Println(isMelkAanwezig && isEiAanwezig && !(isKaasAanwezig))
	// Print of er zowel melk of kaas als ei of boter aanwezig is (true):
	fmt.Println(isMelkAanwezig || isKaasAanwezig || isEiAanwezig || isBoterAanwezig)
	// Print of er geen boter aanwezig is (false):
	fmt.Println(!(isBoterAanwezig))
}

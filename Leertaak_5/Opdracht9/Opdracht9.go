package main

import (
	"fmt"
)

// Data declaration part:

var woordenboekMap map[string]string = map[string]string{"kwakkelplant": "(de, -en) zieke plant of plant die in zijn ontwikkeling achterblijft", "levensloopstress": "(de, g.mv.) stress die werknemers kunnen ervaren als gevolg van gebeurtenissen die zich gedurende alle fasen van hun werkzame leven in hun privésituatie kunnen voordoen.", "floccinaucinihilipilificatie": "(de, g.mv.) gewoonte om handelingen of uitspraken onbeduidend te vinden of te doen schijnen", "snuffelsoftware": "(de, g.mv.) software die geplaatst wordt op computers, smartphones e.d. om informatie over de gebruiker ervan op te sporen", "teslasocialisme": "(het, g.mv.) (ongunstig) elitaire maar zich als socialisme typerende politieke stroming die niet zozeer de verwezenlijking van de traditionele sociale idealen (spreiding van de welvaart, volksverheffing) nastreeft, als wel beoogt de samenleving duurzaam en klimaatneutraal te maken, zelfs al gaat dat ten koste van sociale rechtvaardigheid en gelijke kansen voor gewone mensen"}

func main() {
	// Print all words and definitions:
	for k, v := range woordenboekMap {
		fmt.Printf("%v: %v\n\n", k, v)
	}
	// Check whether floccinaucinihillipilificatie exists in woordenboekMap:
	fmt.Println("Does floccinaucinihillipilificatie exist?")
	_, ok := woordenboekMap["floccinaucinihillipilificatie"]
	fmt.Println(ok)
	// Check whether snuffelsoftware exists in woordenboekMap:
	fmt.Println("Does snuffelsoftware exist?")
	_, ok = woordenboekMap["snuffelsoftware"]
	fmt.Println(ok)
	// Check whether toetsenbordterriër exists in woordenboekMap:
	fmt.Println("Does toetsenbordterriër exist?")
	_, ok = woordenboekMap["toetsenbordterriër"]
	fmt.Println(ok)
	// Check whether kwakkelplant exists in woordenboekMap:
	fmt.Println("Does kwakkelplant exist?")
	_, ok = woordenboekMap["kwakkelplant"]
	fmt.Println(ok)
	// Delete floccinaucinihilipilificatie from the map:
	delete(woordenboekMap, "floccinaucinihilipilificatie")
}

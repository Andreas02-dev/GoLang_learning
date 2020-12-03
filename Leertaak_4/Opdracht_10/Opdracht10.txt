package main

import (
	"fmt"
)

// Data declaration part:

var (
	reizigers           []string           = []string{"Reiziger1", "Reiziger2"}
	balans              map[string]float64 = map[string]float64{"Reiziger1": 100.00, "Reiziger2": 80.00}
	reizigerReisproduct map[string]string  = map[string]string{"Reiziger1": "40% Korting", "Reiziger2": "20% Korting"}
	kostenreis          map[string]float64 = map[string]float64{"Standaardreis": 14.30}
	// Instaptarief in euro's:
	instaptarief = 20.00
)

func main() {
	// Add 2 new travelers and travel routes and give them the Standard reisproduct:
	reizigers = append(reizigers, "Reiziger3", "Reiziger4")
	balans["Reiziger3"] = 200.00
	balans["Reiziger4"] = 300.00
	kostenreis["Utrecht-->Ede-Wageningen"] = 16.24
	kostenreis["Nieuwegein-->Utrecht"] = 6.76
	reizigerReisproduct["Reiziger3"] = "Standaard"
	reizigerReisproduct["Reiziger4"] = "Standaard"
	// Print the balance of the 2 new travelers:
	saldo("Reiziger3")
	saldo("Reiziger4")
	// Have the 2 travelers check in:
	inchecken("Reiziger3", instaptarief)
	inchecken("Reiziger4", instaptarief)
	// Print the balance of the 2 new travelers:
	saldo("Reiziger3")
	saldo("Reiziger4")
	// Have the 2 travelers check out for the 2 different travel routes::
	uitchecken("Reiziger3", "Nieuwegein-->Utrecht", instaptarief)
	uitchecken("Reiziger4", "Utrecht-->Ede-Wageningen", instaptarief)
	// Print the balance of the 2 new travelers:
	saldo("Reiziger3")
	saldo("Reiziger4")
}

// saldo geeft de huidige balans weer van een reiziger in de standaard uitvoer:
func saldo(reiziger string) {
	fmt.Println(balans[reiziger])
}

// Laat de reiziger inchecken door het instaptarief van zijn balans te halen:
func inchecken(reiziger string, instaptarief float64) {
	balans[reiziger] = balans[reiziger] - instaptarief
}

// Laat de reiziger uitchecken door de gemaakte kosten te verrekenen en daarbij rekening te houden met het reizigersproduct.
// Er wordt gebruik gemaakt van een switch statement om te bepalen welke kosten moeten worden gerekend.
func uitchecken(reiziger, reis string, instaptarief float64) {
	var aanpasBedrag float64
	// Check which reisproduct the traveler has and change the balance as balance - travel cost:
	switch reizigerReisproduct[reiziger] {
	case "Standaard":
		aanpasBedrag = kostenreis[reis] - instaptarief
		balans[reiziger] = balans[reiziger] - aanpasBedrag
	case "20% Korting":
		aanpasBedrag = (kostenreis[reis] / 100 * 80) - instaptarief
		balans[reiziger] = balans[reiziger] - aanpasBedrag
	case "40% Korting":
		aanpasBedrag = (kostenreis[reis] / 100 * 60) - instaptarief
		balans[reiziger] = balans[reiziger] - aanpasBedrag
	case "Altijd Vrij":
		balans[reiziger] = balans[reiziger] - aanpasBedrag
	default:
		// Reisproduct unknown:
		fmt.Println("Onbekend reisproduct, transactie mislukt.")
	}
}

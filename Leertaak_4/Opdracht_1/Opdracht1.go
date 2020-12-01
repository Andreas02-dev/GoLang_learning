package main

import (
	"fmt"
)

// Data declaration part:

var (
	// De balans van reiziger 1 in euro's:
	balance1 = 100.00
	// De balans van reiziger 2 in euro's:
	balance2 = 80.00
	// Het reisproduct van reiziger 1:
	reisproduct1 = "40% Korting"
	// Het reisproduct van reiziger 2:
	reisproduct2 = "20% Korting"
	// Instaptarief in euro's:
	instaptarief = 20.00
	// Kosten van de 1ste reis in euro's
	kostenreis = 14.30
)

func main() {
	// Reiziger 1 checkt in en reist:
	balance1 = inchecken(balance1, instaptarief)
	saldo(1)
	// Reiziger 1 checkt uit:
	balance1 = uitchecken(reisproduct1, balance1, kostenreis, instaptarief)
	saldo(1)
	// Reiziger 2 checkt in en reist:
	balance2 = inchecken(balance2, instaptarief)
	saldo(2)
	// Reiziger 2 checkt uit:
	balance2 = uitchecken(reisproduct2, balance2, kostenreis, instaptarief)
	saldo(2)
	// Reiziger 2 verandert het reisproduct naar "Altijd Vrij"
	reisproduct2 = "Altijd Vrij"
	// Reiziger 2 checkt in en reist
	balance2 = inchecken(balance2, instaptarief)
	saldo(2)
}

// saldo geeft de huidige balans weer van een reiziger in de standaard uitvoer:
func saldo(reiziger int) { // Balance float64 is weggehaald omdat deze volstrekt onnodig is.
	switch reiziger {
	case 1:
		fmt.Println(balance1)
	case 2:
		fmt.Println(balance2)
	}
}

// Laat de reiziger inchecken door het instaptarief van zijn balans te halen.
// De functie levert een float64 op die de nieuwe balans voor de reiziger bevat.
func inchecken(balance, instaptarief float64) float64 {
	return balance - instaptarief
}

// Laat de reiziger uitchecken door de gemaakte kosten te verrekenen en daarbij rekening te houden met het reizigersproduct.
// Er wordt gebruik gemaakt van een switch statement om te bepalen welke kosten moeten worden gerekend.
// De functie levert een float64 op die de nieuwe balans voor de reiziger bevat.
func uitchecken(product string, balance, voltarief, instaptarief float64) float64 {
	var aanpasBedrag float64
	switch product {
	case "Standaard":
		aanpasBedrag = voltarief - instaptarief
		balance = balance - aanpasBedrag
		return balance
	case "20% Korting":
		aanpasBedrag = (voltarief / 100 * 80) - instaptarief
		balance = balance - aanpasBedrag
		return balance
	case "40% Korting":
		aanpasBedrag = (voltarief / 100 * 60) - instaptarief
		balance = balance - aanpasBedrag
		return balance
	case "Altijd Vrij":
		balance = balance - aanpasBedrag
		return balance
	default:
		fmt.Println("Onbekende reiziger.")
		return balance
	}
}

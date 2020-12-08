package main

// Data declaration part:

var (
	hoofdstedenNederlandArray [12]string        = [12]string{"Maastricht", "Middelburg", "Den Bosch", "Arnhem", "Utrecht", "Den Haag", "Haarlem", "Lelystad", "Zwolle", "Leeuwarden", "Groningen", "Assen"}
	hoofdstedenNederlandSlice []string          = []string{"Maastricht", "Middelburg", "Den Bosch", "Arnhem", "Utrecht", "Den Haag", "Haarlem", "Lelystad", "Zwolle", "Leeuwarden", "Groningen", "Assen"}
	hoofdstedenNederlandMap   map[int]string    = map[int]string{1: "Maastricht", 2: "Middelburg", 3: "Den Bosch", 4: "Arnhem", 5: "Utrecht", 6: "Den Haag", 7: "Haarlem", 8: "Lelystad", 9: "Zwolle", 10: "Leeuwarden", 11: "Groningen", 12: "Assen"}
	hoofdstedenNederlandMap2  map[string]string = map[string]string{"Limburg (Nederland)": "Maastricht", "Zeeland": "Middelburg", "Noord-Brabant": "Den Bosch", "Gelderland": "Arnhem", "Utrecht": "Utrecht", "Zuid-Holland": "Den Haag", "Noord-Holland": "Haarlem", "Flevoland": "Lelystad", "Overijssel": "Zwolle", "Friesland": "Leeuwarden", "Groningen": "Groningen", "Drenthe": "Assen", "Antwerpen": "Antwerpen", "Limburg": "Hasselt", "Oost-Vlaanderen": "Gent", "Vlaams-Brabant": "Leuven", "West-Vlaanderen": "Brugge", "Henegouwen": "Bergen", "Luik": "Luik", "Luxemburg": "Aarlen", "Namen": "Namen", "Waals-Brabant": "Waver"}
)

func main() {
	// Add the provinces of Belgium to the hoofdstedenNederlandMap:
	hoofdstedenNederlandMap[13] = "Antwerpen"
	hoofdstedenNederlandMap[14] = "Hasselt"
	hoofdstedenNederlandMap[15] = "Gent"
	hoofdstedenNederlandMap[16] = "Leuven"
	hoofdstedenNederlandMap[17] = "Brugge"
	hoofdstedenNederlandMap[18] = "Bergen"
	hoofdstedenNederlandMap[19] = "Luik"
	hoofdstedenNederlandMap[20] = "Aarlen"
	hoofdstedenNederlandMap[21] = "Namen"
	hoofdstedenNederlandMap[22] = "Waver"
	// Delete the capital Leeuwarden from the hoofdstedenNederlandMap:
	delete(hoofdstedenNederlandMap, 10)
	// hoofdstedenNederlandMap2 has the provinces and capitals per province of both the Netherlands and Belgium.
}

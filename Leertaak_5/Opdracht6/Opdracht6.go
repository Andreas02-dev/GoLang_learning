package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Data declaration part:

// UsefulData is the minimalised struct:
type UsefulData []struct {
	Country string `json:"country"`
	Wins    int    `json:"wins"`
	Draws   int    `json:"draws"`
	Losses  int    `json:"losses"`
	Points  int    `json:"points"`
}

func main() {
	// Create a variable type UsefulData:
	DataStruct := UsefulData{}
	DataStruct2 := UsefulData{}
	// Retreieve the json byteslice:
	bs := retrieveDataFromFile("Lesweek5_LT5.3_Opdracht6_JSON-start.json")
	// Unmarshal it into the DataStruct:
	err := json.Unmarshal(bs, &DataStruct)
	if err != nil {
		fmt.Println(err)
	}
	// Print the DataStruct:
	fmt.Println(DataStruct)
	// Marshal it into json:
	bs2, err := json.Marshal(DataStruct)
	if err != nil {
		fmt.Println(err)
	}
	// Write it to a new json file:
	writeFile("NewJSONfromStruct.json", bs2)
	// Read the new json file:
	bs3 := retrieveDataFromFile("NewJSONfromStruct.json")
	// Unmarshal it into a new UsefulData typed variable:
	err = json.Unmarshal(bs3, &DataStruct2)
	if err != nil {
		fmt.Println(err)
	}
	// Print the DataStruct2:
	fmt.Println(DataStruct2)
}

// retrieveDataFromFile retrieves the body of a file and returns the byteslice:
func retrieveDataFromFile(jsonFileName string) []byte {
	// Open our jsonFile:
	jsonFile, err := os.Open(jsonFileName)
	// if os.Open returns an error then handle it:
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	body, err2 := ioutil.ReadAll(jsonFile)
	if err2 != nil {
		fmt.Println(err2)
	}
	return body
}

// writeFile takes a filename and a byte slice and writes it to the file.
// If an error is thrown, it will print it to the terminal:
func writeFile(jsonFileName string, data []byte) {
	err := ioutil.WriteFile(jsonFileName, data, 0655)
	if err != nil {
		fmt.Println(err)
	}
}

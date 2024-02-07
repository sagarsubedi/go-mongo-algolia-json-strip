package main

import (
	"consultancyToAlgolia/model"
	"log"
	"os"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	jsonFile := "./updated_consultancies.json"

	byteValue, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	// ? Deserialize json data
	var consultanciesInput []model.ConsultancyInput
	err = json.Unmarshal(byteValue, &consultanciesInput)

	if err != nil {
		log.Fatal(err)
	}

	consultancies := make([]model.ConsultancyForAlgolia, len(consultanciesInput))
	for i, c := range consultanciesInput {
		consultancies[i] = model.ConsultancyForAlgolia{
			ID:      c.ID.OID,
			Name:    c.Name,
			Address: c.Address,
			City:    c.City,
		}
	}

	filteredData, err := json.MarshalIndent(consultancies, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("consultancies_for_algolia.json", filteredData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Filtered data written to filtered_consultancies.json")
}

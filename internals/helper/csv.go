package helper

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/masrayfa/internals/models/domain"
)

func GenerateCSV(data []domain.Channel) (string, error) {
	file, err := os.Create("data.csv")
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"time", "value 1", "value 2", "value 3", "value 4", "value 5", "value 6", "value 7", "value 8", "value 9", "value 10"})

	// Write CSV records
	for _, d := range data {
		var record []interface{}
		record = append(record, d.Time)
		for _, v := range d.Value{
			record = append(record, v)
		}

		var stringRecord []string
		for _, v := range record {
			stringRecord = append(stringRecord, fmt.Sprintf("%v", v))
		}
		
		if err := writer.Write(stringRecord); err != nil {
			log.Println("Error writing record to csv:", err)
		}
	}

	// for _, row := range data {
	// 	for _, v := range row.Value {
	// 		fmt.Printf("value data csv: %v, ", v)
	// 	}
	// }

	return file.Name(), nil
}
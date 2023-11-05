package read_csv

import (
	"encoding/csv"
	"os"
)

func ReadCsvFile(file string, numQuestions *int) ([][]string, error) {
	var records [][]string

	// open csv file
	f, err := os.Open(file)

	if err != nil {
		return records, err
	}
	defer f.Close()

	//initialize new csv reader
	csvReader := csv.NewReader(f)

	records, err = csvReader.ReadAll()
	if err != nil {
		return records, err
	}
	*numQuestions = len(records)
	return records, nil
}

package parser

import (
	"encoding/csv"
	"os"
)

func ReadCSV(path string) ([][]string, error){
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	return reader.ReadAll()
}
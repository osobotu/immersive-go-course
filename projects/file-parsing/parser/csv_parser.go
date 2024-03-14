package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type csvParser struct {
	FileName string
}

func (c csvParser) Parse() ([]Player, error) {
	path := GetPath("data.csv")

	file, err := os.Open(path)
	if err != nil {
		err := fmt.Errorf("could not open file: %v", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()
	if err != nil {
		err := fmt.Errorf("could not read csv data: %v", err)
		return nil, err
	}

	players := make([]Player, 0)
	for _, row := range data[1:] {
		score, err := strconv.Atoi(row[1])
		if err != nil {
			err := fmt.Errorf("cannot convert string score to int: %v", err)
			return nil, err
		}
		player := Player{
			Name:  row[0],
			Score: score,
		}
		players = append(players, player)
	}
	return players, nil
}

func (c csvParser) String() string {
	return "CSV Parser"
}

func NewCSVParser() FileParser {
	return csvParser{}
}

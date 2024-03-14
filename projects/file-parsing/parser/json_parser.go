package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type jsonParser struct {
}

func (j jsonParser) Parse() ([]Player, error) {
	path := GetPath("json.txt")

	file, err := os.Open(path)
	if err != nil {
		err := fmt.Errorf("could not open file: %v", err)
		return nil, err
	}
	defer file.Close()

	byteFile, _ := io.ReadAll(file)
	players := make([]Player, 0)

	json.Unmarshal(byteFile, &players)

	return players, nil
}

func NewJSONParser() FileParser {
	return jsonParser{}
}

func (c jsonParser) String() string {
	return "JSON Parser"
}

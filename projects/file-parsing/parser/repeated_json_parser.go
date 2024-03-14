package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type repeatedJSONParser struct {
}

func NewRepeatedJSONParser() FileParser {
	return repeatedJSONParser{}
}

func (c repeatedJSONParser) String() string {
	return "Repeated JSON Parser"
}

func (rj repeatedJSONParser) Parse() ([]Player, error) {
	path := GetPath("repeated-json.txt")

	file, err := os.Open(path)
	if err != nil {
		err := fmt.Errorf("could not open file: %v", err)
		return nil, err
	}
	defer file.Close()

	players := make([]Player, 0)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		// 35 is the byte code for #
		if !(fileScanner.Text()[0] == 35) {
			var player Player
			json.Unmarshal([]byte(fileScanner.Text()), &player)
			players = append(players, player)
		}
	}

	return players, nil
}

package main

import (
	"file_parsing/parser"
	"fmt"
	"os"
)

func main() {
	parsers := make([]parser.FileParser, 0)

	csvParser := parser.NewCSVParser()
	parsers = append(parsers, csvParser)

	jsonParser := parser.NewJSONParser()
	parsers = append(parsers, jsonParser)

	repeatedJSONParser := parser.NewRepeatedJSONParser()
	parsers = append(parsers, repeatedJSONParser)

	binaryParser := parser.NewBinaryParser()
	parsers = append(parsers, binaryParser)

	for _, item := range parsers {
		players, err := item.Parse()
		if err != nil {
			err := fmt.Errorf("could not parse file using parser: %v", err)
			fmt.Fprint(os.Stderr, err)
		}
		if len(players) == 0 {
			return
		}
		max, min := parser.FindMaxAndMin(players)
		fmt.Fprintln(os.Stdout, "Parser:", item)
		fmt.Fprintln(os.Stdout, "Highest Score:", max)
		fmt.Fprintln(os.Stdout, "Lowest Score:", min)
		fmt.Println()
	}
}

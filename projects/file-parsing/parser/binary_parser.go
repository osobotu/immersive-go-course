package parser

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type binaryParser struct {
}

func NewBinaryParser() FileParser {
	return binaryParser{}
}

func (b binaryParser) String() string {
	return "Binary Parser"
}

func (b binaryParser) Parse() ([]Player, error) {
	path := GetPath("custom-binary-le.bin")
	// path := GetPath("custom-binary-be.bin")

	// open the file
	file, err := os.Open(path)
	if err != nil {
		err := fmt.Errorf("could not open file: %v", err)
		return nil, err
	}
	defer file.Close()

	// read the first two bytes to determine endianness
	var endianness binary.ByteOrder
	bytesOrderArray := make([]byte, 2)
	n, err := file.Read(bytesOrderArray)
	if n != 2 {
		err := fmt.Errorf("should read 2 bytes not %v bytes", n)
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	if bytesOrderArray[0] == 0xFE && bytesOrderArray[1] == 0xFF {
		endianness = binary.BigEndian
	} else if bytesOrderArray[0] == 0xFF && bytesOrderArray[1] == 0xFE {
		endianness = binary.LittleEndian
	} else {
		err := fmt.Errorf("unknown endianness for binary format: %v", bytesOrderArray)
		return nil, err
	}

	// Parse records
	players := make([]Player, 0)
	for {
		// Read the score (4 bytes)
		var score int32
		err := binary.Read(file, endianness, &score)
		if err != nil {
			if err == io.EOF {
				break // Reached end of file
			}
			fmt.Println("Error reading score:", err)
			break
		}

		// Read the name of the player (until null terminator)
		name := ""
		for {
			var char byte
			err := binary.Read(file, endianness, &char)
			if err != nil {
				fmt.Println("Error reading name character:", err)
				return nil, err
			}
			if char == 0 {
				break
			}
			name += string(char)
		}
		players = append(players, Player{Name: name, Score: int(score)})
	}

	return players, nil
}

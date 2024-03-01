package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Execute() {
	// get current working directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if !(len(os.Args) > 1) {
		fmt.Fprintln(os.Stderr, "Usage: go-cat <file-name>")
		return
	}

	start := 1
	showLineNumber := os.Args[1] == "-n"

	if showLineNumber {
		// start after the modifier
		start = 2
	}

	files := os.Args[start:]

	for _, file := range files {
		path := filepath.Join(pwd, file)
		fs, err := os.Stat(path)
		if err != nil {
			err = fmt.Errorf("could not get file stats: %v", err)
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if fs.IsDir() {
			fmt.Fprintf(os.Stderr, "go-cat: %s is a directory\n", fs.Name())
			return
		}
		if showLineNumber {
			readAndPrintLineByLine(path)
		} else {
			readAndPrintWholeContent(path)
		}

	}
}

func readAndPrintLineByLine(path string) {
	readFile, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fileScanner := bufio.NewScanner(readFile)
	count := 1
	for fileScanner.Scan() {
		fmt.Fprintln(os.Stdout, count, fileScanner.Text())
		count++
	}

	readFile.Close()
}

func readAndPrintWholeContent(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("could not read file content: %v", err)
		fmt.Fprintln(os.Stderr, err)
	}
	os.Stdout.Write(data)
}

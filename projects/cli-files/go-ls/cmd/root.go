package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func Execute() {
	// get current working directory (cwd)
	path, err := os.Getwd()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	// get folder
	if len(os.Args) > 1 {
		firstArg := os.Args[1]
		if firstArg == "-h" {
			fmt.Println("go-ls is a go implementation of the ls tool for listing files and directories")
			return
		}
		path = filepath.Join(path, firstArg)
	}

	fs, err := os.Stat(path)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	if fs.IsDir() {
		items, err := os.ReadDir(path)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		// list directories and files in cwd
		for _, file := range items {
			fmt.Fprint(os.Stdout, file.Name(), "  ")
			return
		}
		fmt.Println()
		return
	}
	fmt.Println(fs.Name())

}

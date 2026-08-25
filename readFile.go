package main

// http://zetcode.com/golang/readfile/

import (
	"bufio"
	"log"
	"os"
)

func readFile(file string) []string {

	var lines []string

	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close %s: %v", file, err)
		}
	}()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

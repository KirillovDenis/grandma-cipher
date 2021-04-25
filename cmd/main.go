package main

import (
	"bufio"
	"grandma-cipher/cmd/generator"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	minWordLength := 3
	dict, err := getDict("cmd/usa.txt", minWordLength)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	gen := generator.NewGenerator(dict)
	result := gen.Recursive(4, 20, 24, minWordLength)
	elapsed := time.Since(start)

	log.Printf("Elapsed: %v\n %v\n", elapsed, result)
}

func getDict(filename string, minWordSize int) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	dict := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) >= minWordSize && isASCII(txt) {
			dict = append(dict, strings.ToLower(txt))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dict, file.Close()
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 65 || (s[i] > 90 && s[i] < 97) || s[i] > 122 {
			return false
		}
	}
	return true
}

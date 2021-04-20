package main

import (
	"bufio"
	"grandma-cipher/cmd/generator"
	"log"
	"os"
	"time"
)

func main() {
	dict, err := getDict("cmd/usa.txt", 3)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	gen := generator.NewGenerator(dict)

	results, err := gen.GreedyMult(4, 20, 24, 9)
	elapsed := time.Since(start)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Elapsed: %v\n %#v\n", elapsed, results)
	}
}

func getDict(filename string, minWordSize int) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	dict := make([]string, 0, 70000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) >= minWordSize {
			dict = append(dict, txt)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dict, file.Close()
}

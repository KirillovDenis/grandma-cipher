package generator

import (
	"fmt"
	"log"
	"sort"
)

type Generator struct {
	dict   []string
	sorted bool
}

type GenResults struct {
	Words  []string
	Length int
	Weight int
}

func NewGenerator(dict []string) *Generator {
	return &Generator{dict, false}
}

func NewGenResults(words []string) (*GenResults, error) {
	weight, err := computeWeights(words)
	if err != nil {
		return nil, err
	}

	return &GenResults{
		Words:  words,
		Length: wordsLength(words),
		Weight: weight,
	}, nil
}

func (g *Generator) sortDict() {
	sort.Slice(g.dict, func(i, j int) bool {
		w1, err := weight(g.dict[i])
		if err != nil {
			log.Println(err)
		}
		w2, err := weight(g.dict[j])
		if err != nil {
			log.Println(err)
		}

		return w1 < w2
	})
	g.sorted = true
}

func (g *Generator) Greedy(size, min, max int) (*GenResults, error) {
	if len(g.dict) < size {
		return nil, fmt.Errorf("dict too small")
	}

	if !g.sorted {
		g.sortDict()
	}

	result := g.dict[:size]

	for _, word := range g.dict {
		current := g.getWords(word, size)
		if ok, err := better(current, result, min, max); err != nil {
			return nil, err
		} else if ok {
			result = getCopy(current)
		}
	}

	return NewGenResults(result)
}

func (g *Generator) GreedyMult(size, min, max, lim int) (*GenResults, error) {
	if len(g.dict) < size {
		return nil, fmt.Errorf("dict too small")
	}

	if !g.sorted {
		g.sortDict()
	}

	result := g.dict[:size]

	for _, word := range g.dict {
		candidates := g.getWordsCandidates(nil, nil, word, size, lim)
		for _, current := range candidates {
			if ok, err := better(current, result, min, max); err != nil {
				return nil, err
			} else if ok {
				result = getCopy(current)
			}
		}
	}

	return NewGenResults(result)
}

func (g *Generator) getWords(startWord string, size int) []string {
	words := []string{startWord}
	for len(words) < size {
		found := false
		lastWord := words[len(words)-1]
		nearestLetters := getNearest(lastWord)

		for _, word := range g.dict {
			if startWith(word, nearestLetters) && !contains(words, word) {
				found = true
				words = append(words, word)
				break
			}
		}

		if !found {
			words = append(words, g.dict[0])
		}
	}

	return words
}

func startWith(word string, asciiLetters []byte) bool {
	for _, letter := range asciiLetters {
		if letter == word[0] {
			return true
		}
	}
	return false
}

func (g *Generator) getWordsCandidates(wordsCandidates [][]string, words []string, lastWord string, size, lim int) [][]string {
	if len(words) >= size {
		return wordsCandidates
	}

	words = append(words, lastWord)
	if len(words) == size {
		return append(wordsCandidates, getCopy(words))
	}

	nearestWords := getNearestWords(g.dict, lastWord)
	if len(nearestWords) > lim {
		nearestWords = nearestWords[:lim]
	}

	for _, word := range nearestWords {
		if !contains(words, word) {
			wordsCandidates = g.getWordsCandidates(wordsCandidates, words, word, size, lim)
		}
	}

	return wordsCandidates
}

func getNearestWords(dict []string, lastWord string) []string {
	var results []string
	nearest := getNearest(lastWord)
	for _, word := range dict {
		if len(nearest) == 0 {
			break
		}
		if index := getIndex(nearest, word[0]); index != -1 {
			results = append(results, word)
			nearest = append(nearest[:index], nearest[index+1:]...)
		}
	}

	return results
}

func contains(array []string, elem string) bool {
	for _, element := range array {
		if elem == element {
			return true
		}
	}

	return false
}

func getIndex(array []byte, elem byte) int {
	for i, element := range array {
		if elem == element {
			return i
		}
	}

	return -1
}

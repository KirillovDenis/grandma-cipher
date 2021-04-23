package generator

import (
	"sort"
)

type wordCost struct {
	cost int
	word string
}

type Generator struct {
	dict      []wordCost
	distances map[spacing]int
	bestCost  int
	bestWords []string
}

type GenResults struct {
	Words  []string
	Length int
	Weight int
}

func NewGenerator(dict []string) *Generator {
	preparedDict := make([]wordCost, len(dict))

	for i := range dict {
		preparedDict[i] = wordCost{weight(dict[i]), dict[i]}
	}

	sort.Slice(preparedDict, func(i, j int) bool {
		return preparedDict[i].cost < preparedDict[j].cost
	})

	return &Generator{dict: preparedDict, distances: prepare(), bestCost: 10000}
}

func NewGenResults(words []string) *GenResults {
	return &GenResults{
		Words:  words,
		Length: computeLength(words),
		Weight: computeWeights(words),
	}
}

func prepare() map[spacing]int {
	distances := make(map[spacing]int)

	for i := 0; i < 3; i++ {
		for j := 0; j < len(keyboard[i]); j++ {
			ch := keyboard[i][j]
			for ii := 0; ii < 3; ii++ {
				for jj := 0; jj < len(keyboard[ii]); jj++ {
					ch2 := keyboard[ii][jj]
					distances[spacing{ch, ch2}] = distanceByte(ch, ch2)
				}
			}
		}
	}

	return distances
}

func (g *Generator) Recursive(size, min, max, minWordLength int) *GenResults {
	if len(g.dict) < size {
		return nil
	}

	maxWordLength := max - (size-1)*minWordLength
	for _, word := range g.dict {
		curCost := word.cost
		curLength := len(word.word)
		if curCost >= g.bestCost {
			break
		}
		if curLength > maxWordLength {
			continue
		}

		g.find([]string{word.word}, maxWordLength, curCost, curLength, min, max, size)
	}

	return NewGenResults(g.bestWords)
}

func (g *Generator) find(words []string, maxWordLength, cost, length, min, max, size int) {
	if len(words) == size {
		if length >= min && length <= max {
			g.bestCost = cost
			g.bestWords = words
		}
		return
	}

	lastWord := words[len(words)-1]
	for _, word := range g.dict {
		newCost := cost + word.cost
		if newCost >= g.bestCost {
			break
		}
		wordLen := len(word.word)
		if wordLen > maxWordLength {
			continue
		}
		newCost += g.distances[spacing{lastWord[len(lastWord)-1], word.word[0]}]
		if newCost >= g.bestCost {
			continue
		}
		newLength := length + wordLen
		if newLength > max || contains(words, word.word) {
			continue
		}

		g.find(append(words, word.word), maxWordLength, newCost, newLength, min, max, size)
	}
}

func contains(array []string, elem string) bool {
	for _, element := range array {
		if elem == element {
			return true
		}
	}

	return false
}

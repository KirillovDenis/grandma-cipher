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
	distances [][]int
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

func prepare() [][]int {
	distances := make([][]int, len(alphabet))

	for i := 0; i < len(alphabet); i++ {
		distances[i] = make([]int, len(alphabet))
		for j := 0; j < len(alphabet); j++ {
			distances[i][j] = distanceByte(alphabet[i], alphabet[j])
		}
	}

	return distances
}

type costLength struct {
	cost   int
	length int
}

func (g *Generator) NoRecursive(size, min, max, minWordLength int) *GenResults {
	if len(g.dict) < size {
		return nil
	}

	maxWordLength := max - (size-1)*minWordLength
	g.bestWords = make([]string, size)

	seen := make([]int, size)
	for i := 0; i < size; i++ {
		seen[i] = -1
	}

	costsLengths := make([]costLength, size)

	for _, word := range g.dict {
		curCost := word.cost
		curLength := len(word.word)
		if curCost >= g.bestCost {
			break
		}
		if curLength > maxWordLength {
			continue
		}

		costsLengths[0] = costLength{curCost, curLength}
		words := []string{word.word}
		for len(words) > 0 {
			lastWord := words[len(words)-1]
			seenBound := seen[len(words)]
			found := false
			for i, word := range g.dict {
				newCost := costsLengths[len(words)-1].cost + word.cost
				if newCost >= g.bestCost {
					break
				}
				if i <= seenBound || len(word.word) > maxWordLength {
					continue
				}
				newCost += g.distances[lastWord[len(lastWord)-1]-'a'][word.word[0]-'a']
				if newCost >= g.bestCost {
					continue
				}
				newLength := costsLengths[len(words)-1].length + len(word.word)
				if newLength > max || contains(words, word.word) {
					continue
				}

				seen[len(words)] = i
				if len(words) == size-1 {
					if newLength >= min && newLength <= max {
						g.bestCost = newCost
						copy(g.bestWords, append(words, word.word))
					}
					continue
				}

				costsLengths[len(words)] = costLength{newCost, newLength}
				words = append(words, word.word)
				found = true
				break
			}
			if !found {
				words = words[:len(words)-1]
				for i := len(words) + 1; i < size; i++ {
					seen[i] = -1
				}
			}
		}
	}

	return NewGenResults(g.bestWords)
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
		newCost += g.distances[lastWord[len(lastWord)-1]-'a'][word.word[0]-'a']
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

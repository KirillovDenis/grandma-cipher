package generator

import (
	"fmt"
	"log"
)

type point struct {
	x, y int
}

var byteToPointMap = map[byte]point{
	'q': {0, 0}, 'w': {1, 0}, 'e': {2, 0}, 'r': {3, 0}, 't': {4, 0}, 'y': {5, 0},
	'u': {6, 0}, 'i': {7, 0}, 'o': {8, 0}, 'p': {9, 0},

	'a': {0, 1}, 's': {1, 1}, 'd': {2, 1}, 'f': {3, 1}, 'g': {4, 1}, 'h': {5, 1},
	'j': {6, 1}, 'k': {7, 1}, 'l': {8, 1},

	'z': {0, 2}, 'x': {1, 2}, 'c': {2, 2}, 'v': {3, 2}, 'b': {4, 2}, 'n': {5, 2},
	'm': {6, 2},
}

var pointToByteMap = map[point]byte{
	{0, 0}: 'q', {1, 0}: 'w', {2, 0}: 'e', {3, 0}: 'r', {4, 0}: 't', {5, 0}: 'y',
	{6, 0}: 'u', {7, 0}: 'i', {8, 0}: 'o', {9, 0}: 'p',

	{0, 1}: 'a', {1, 1}: 's', {2, 1}: 'd', {3, 1}: 'f', {4, 1}: 'g', {5, 1}: 'h',
	{6, 1}: 'j', {7, 1}: 'k', {8, 1}: 'l',

	{0, 2}: 'z', {1, 2}: 'x', {2, 2}: 'c', {3, 2}: 'v', {4, 2}: 'b', {5, 2}: 'n',
	{6, 2}: 'm',
}

func getNearest(word string) []byte {
	result := make([]byte, 0, 9)
	if len(word) == 0 {
		for k := range byteToPointMap {
			result = append(result, k)
		}
	}

	endPoint, ok := byteToPointMap[word[len(word)-1]]
	if !ok {
		log.Printf("unsupported ascii char '%d'", word[len(word)-1])
	}

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			newPoint := point{endPoint.x + i, endPoint.y + j}
			char, ok := pointToByteMap[newPoint]
			if ok {
				result = append(result, char)
			}
		}
	}

	return result
}

func better(first, second []string, min, max int) (bool, error) {
	weight1, err := computeWeights(first)
	if err != nil {
		return false, err
	}
	length1 := wordsLength(first)
	if length1 < min || length1 > max {
		return false, nil
	}

	weight2, err := computeWeights(second)
	if err != nil {
		return false, err
	}
	length2 := wordsLength(second)
	if length2 < min || length2 > max {
		return true, nil
	}

	return weight1 < weight2, nil
}

func computeWeights(array []string) (int, error) {
	if len(array) == 0 {
		return -1, fmt.Errorf("array must be not empty")
	}

	last := array[0]
	w, err := weight(last)
	if err != nil {
		return -1, err
	}
	sum := w

	for i := 1; i < len(array); i++ {
		w, err := weight(array[i])
		if err != nil {
			return -1, err
		}
		sum += w
		d, err := distance(last, array[i])
		if err != nil {
			return -1, err
		}
		sum += d
		last = array[i]
	}

	return sum, nil
}

func wordsLength(arr []string) int {
	sumLetters := 0

	for _, word := range arr {
		sumLetters += len(word)
	}

	return sumLetters
}

func getCopy(arr []string) []string {
	result := make([]string, 0, len(arr))
	for _, word := range arr {
		result = append(result, word)
	}

	return result
}

func weight(word string) (int, error) {
	if len(word) == 0 {
		return -1, fmt.Errorf("word must not be empty")
	}

	sum := 0
	last := word[0]
	for i := 1; i < len(word); i++ {
		dist, err := distanceByte(last, word[i])
		if err != nil {
			return -1, err
		}
		last = word[i]
		sum += dist
	}

	return sum, nil
}

func distance(first, second string) (int, error) {
	if len(first) == 0 || len(second) == 0 {
		return -1, fmt.Errorf("slices must not be empty")
	}

	endByte := first[len(first)-1]
	startByte := second[0]

	return distanceByte(endByte, startByte)
}

func distanceByte(first, second byte) (int, error) {
	point1, ok := byteToPointMap[first]
	if !ok {
		return -1, fmt.Errorf("unsupported byte value  %d, allowed only ASCII letters", first)
	}

	point2, ok := byteToPointMap[second]
	if !ok {
		return -1, fmt.Errorf("unsupported byte value %d, allowed only ASCII letters ", second)
	}

	return taxicabDistance(point1, point2), nil
}

func taxicabDistance(first, second point) int {
	return abs(first.x-second.x) + abs(first.y-second.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

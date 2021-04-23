package generator

type spacing struct {
	chEnd, chStart byte
}

var keyboard = [3]string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}

func computeWeights(array []string) int {
	if len(array) == 0 {
		return 0
	}

	last := array[0]
	sum := weight(last)

	for i := 1; i < len(array); i++ {
		sum += weight(array[i])
		sum += distance(last, array[i])
		last = array[i]
	}

	return sum
}

func computeLength(arr []string) int {
	sumLetters := 0
	for _, word := range arr {
		sumLetters += len(word)
	}
	return sumLetters
}

func weight(word string) int {
	sum := 0
	last := word[0]
	for i := 1; i < len(word); i++ {
		sum += distanceByte(last, word[i])
		last = word[i]
	}

	return sum
}

func distance(first, second string) int {
	if len(first) == 0 || len(second) == 0 {
		return 0
	}

	endByte := first[len(first)-1]
	startByte := second[0]

	return distanceByte(endByte, startByte)
}

func distanceByte(first, second byte) int {
	var x1, y1, x2, y2 int
	for i := 0; i < 3; i++ {
		for j := 0; j < len(keyboard[i]); j++ {
			if keyboard[i][j] == first {
				x1, y1 = i, j
			}
			if keyboard[i][j] == second {
				x2, y2 = i, j
			}
		}
	}

	return taxicabDistance(x1, y1, x2, y2)
}

func taxicabDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

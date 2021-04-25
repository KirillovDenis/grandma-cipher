package generator

var keyboard = [3]string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
var alphabet = "abcdefghijklmnopqrstuvwxyz"

func weight(word string) int {
	sum := 0
	last := word[0]
	for i := 1; i < len(word); i++ {
		sum += distanceByte(last, word[i])
		last = word[i]
	}

	return sum
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

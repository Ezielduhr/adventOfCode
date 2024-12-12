package ceresSearch

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type wordPuzzle struct {
	letters         [][]string
	possibleStrings []string
	amountOfMatches int
	amountOfF       int
}

func readFromFile(filePath string) ([][]string, error) {
	var sArray [][]string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var currentSlice []string
		for _, char := range scanner.Text() {
			currentSlice = append(currentSlice, string(char))
		}
		sArray = append(sArray, currentSlice)
	}

	return sArray, nil
}

func (wp *wordPuzzle) calculateAllTheThings() error {
	yMaxLength := len(wp.letters)
	xMaxLength := len(wp.letters[0])

	for y, _ := range wp.letters {
		horizontal := ""
		for _, yValue := range wp.letters[y] {
			horizontal = horizontal + yValue
		}
		wp.possibleStrings = append(wp.possibleStrings, horizontal)
	}
	for x := range len(wp.letters) {
		vertical := ""
		for y, _ := range wp.letters[x] {
			vertical = vertical + wp.letters[y][x]
		}
		wp.possibleStrings = append(wp.possibleStrings, vertical)
	}
	for offset, _ := range wp.letters[0] {
		leadingDiagonalHorizontal := ""
		for y, _ := range wp.letters {
			for x, _ := range wp.letters[0] {
				if y+offset == x {
					if y+offset < yMaxLength && x < xMaxLength {
						leadingDiagonalHorizontal = leadingDiagonalHorizontal + wp.letters[y][x]
					}
				}
			}

		}
		if offset != 0 {
			// first diagonal we add during the leadingDiagonalVertical
			wp.possibleStrings = append(wp.possibleStrings, leadingDiagonalHorizontal)
		}
		leadingDiagonalVertical := ""
		for y, _ := range wp.letters {
			for x, _ := range wp.letters[0] {
				if y == x+offset {
					if y < yMaxLength && x+offset < xMaxLength {
						leadingDiagonalVertical = leadingDiagonalVertical + wp.letters[y][x]
					}
				}
			}
		}
		wp.possibleStrings = append(wp.possibleStrings, leadingDiagonalVertical)
	}
	for offset, _ := range wp.letters[0] {
		antiDiagonalHorizontal := ""
		for y, _ := range wp.letters {
			for x, _ := range wp.letters[0] {
				x = len(wp.letters[0]) - (x + 1)
				if y+offset == len(wp.letters[0])-(x+1) {
					if y+offset < yMaxLength && x >= 0 {
						antiDiagonalHorizontal = antiDiagonalHorizontal + wp.letters[y][x]
					}
				}
			}

		}
		if offset != 0 {
			// first diagonal we add during the antiDiagonalVertical
			wp.possibleStrings = append(wp.possibleStrings, antiDiagonalHorizontal)
		}
		antiDiagonalVertical := ""
		for y, _ := range wp.letters {
			for x, _ := range wp.letters[0] {
				x = len(wp.letters[0]) - (x + 1)
				if y == len(wp.letters[0])-(x+1)+offset {
					if y >= 0 && x <= xMaxLength {
						antiDiagonalVertical = antiDiagonalVertical + wp.letters[y][x]
					}
				}
			}
		}
		wp.possibleStrings = append(wp.possibleStrings, antiDiagonalVertical)
	}
	return nil
}

func (wp *wordPuzzle) checkAllTheThings(regexp *regexp.Regexp) (int, error) {
	for _, sToCheck := range wp.possibleStrings {
		foundStrings := regexp.FindAllString(sToCheck, -1)
		wp.amountOfMatches += len(foundStrings)
	}

	return wp.amountOfMatches, nil
}

func (wp *wordPuzzle) getValue(y int, x int) string {
	if y < 0 || x < 0 || y >= len(wp.letters) || x >= len(wp.letters) {
		return ""
	}
	return wp.letters[y][x]
}

func (wp *wordPuzzle) xMasFEry() error {
	forward := "MAS"
	reverse := "SAM"
	matches := 0

	for y, _ := range wp.letters {
		for x, _ := range wp.letters[y] {
			topLeft := wp.getValue(y-1, x-1)
			topRight := wp.getValue(y-1, x+1)
			current := wp.getValue(y, x)
			bottomLeft := wp.getValue(y+1, x-1)
			bottomRight := wp.getValue(y+1, x+1)

			leadingDiagonal := topLeft + current + bottomRight
			antiDiagonal := topRight + current + bottomLeft

			if (leadingDiagonal == forward || leadingDiagonal == reverse) && (antiDiagonal == forward || antiDiagonal == reverse) {
				matches++
			}
		}
	}

	wp.amountOfF = matches
	return nil
}

func Main() {
	wp := wordPuzzle{}
	wp.letters, _ = readFromFile("twentyFour/ceresSearch/ceresSearch.txt")
	_ = wp.calculateAllTheThings()

	xmasRegex, _ := regexp.Compile("XMAS")
	_, _ = wp.checkAllTheThings(xmasRegex)
	smaxRegex, _ := regexp.Compile("SAMX")
	_, _ = wp.checkAllTheThings(smaxRegex)
	fmt.Println(wp.amountOfMatches)

	_ = wp.xMasFEry()
	fmt.Println(wp.amountOfF)
}

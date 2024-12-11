package mullItOver

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MulReader struct {
	rawInput            []string
	multipleCations     []string
	sumOfMultiplication int
}

func readFromFile(filePath string) ([]string, error) {
	var multiplications []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		multiplications = append(multiplications, scanner.Text())
	}

	return multiplications, nil
}
func parseStringForMul(text string, regex *regexp.Regexp) []string {
	var multiplications []string
	multiplications = regex.FindAllString(text, -1)
	return multiplications
}

func (mr *MulReader) calculateMultiplication() error {
	enabled := true

	for _, multipleCation := range mr.multipleCations {
		if multipleCation == "don't()" {
			enabled = false
		} else if multipleCation == "do()" {
			enabled = true
		}

		if multipleCation[:3] == "mul" && enabled {
			splitStrings := strings.Split(multipleCation, ",")
			first, _ := strconv.Atoi(splitStrings[0][4:])
			second, _ := strconv.Atoi(splitStrings[1][:len(splitStrings[1])-1])
			mr.sumOfMultiplication += first * second
		}
	}

	return nil
}

func Main() {
	mr := MulReader{}
	mr.rawInput, _ = readFromFile("twentyFour/mullItOver/mullItOver.txt")
	regex, _ := regexp.Compile("do\\(\\)|don't\\(\\)|mul(\\(\\d{1,3},\\d{1,3}\\))")
	for _, line := range mr.rawInput {
		mr.multipleCations = append(mr.multipleCations, parseStringForMul(line, regex)...)
	}
	_ = mr.calculateMultiplication()

	fmt.Println(mr.sumOfMultiplication)

}

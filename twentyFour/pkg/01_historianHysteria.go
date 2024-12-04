package historianHysteria

import (
	"bufio"
	"iter"
	"log"
	"maps"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Zip[T, U any](t []T, u []U) iter.Seq2[T, U] {
	return func(yield func(T, U) bool) {
		for i := range min(len(t), len(u)) {
			if !yield(t[i], u[i]) {
				return
			}
		}
	}
}

type ListCompare struct {
	FirstList     []int
	SecondList    []int
	totalDistance int
}

func (lc *ListCompare) CalculateDistance() {
	lc.totalDistance = 0
	sort.Ints(lc.FirstList)
	sort.Ints(lc.SecondList)
	// Go why u no zip ??
	// https://stackoverflow.com/questions/26957040/how-to-combine-slices-into-a-slice-of-tuples-in-go-implementing-python-zip-fu
	zipList := maps.Collect(Zip(lc.FirstList, lc.SecondList))

	for first, second := range zipList {
		result := first - second
		lc.totalDistance += int(math.Abs(float64(result)))
	}

}

func (lc *ListCompare) readFromFile(filePath string) {
	firstSlice := []int{0}
	secondSlice := []int{0}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scan1 := bufio.NewScanner(file)
	for scan1.Scan() {
		first, second, _ := strings.Cut(scan1.Text(), " ")

		iFirst, _ := strconv.Atoi(first)
		firstSlice = append(firstSlice, iFirst)

		iSecond, _ := strconv.Atoi(strings.TrimSpace(second))
		secondSlice = append(secondSlice, iSecond)

	}
	lc.FirstList = firstSlice
	lc.SecondList = secondSlice
}

func Main() {
	listCompare := new(ListCompare)
	listCompare.readFromFile("twentyFour/resources/01_historianHysteria.txt")
	listCompare.CalculateDistance()
	println(listCompare.totalDistance)
}

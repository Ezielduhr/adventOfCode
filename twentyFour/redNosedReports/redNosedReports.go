package redNosedReports

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type SafetyReport struct {
	levels          []int
	diffs           []int
	safe            bool
	ProblemDampener bool
}

func calculateDiff(levels []int) ([]int, error) {
	var diffs []int

	previousLevel := 0
	diff := 0
	for i, level := range levels {
		if i >= 1 {
			diff = level - previousLevel
			diffs = append(diffs, diff)
		}
		previousLevel = level
		diff = 0
	}
	if len(diffs) != len(levels)-1 {
		return nil, fmt.Errorf("expected diffs length to be %d got %d", len(levels)-1, diffs)
	}

	return diffs, nil
}

func (sr *SafetyReport) SetSafetyLevels(levels []int) error {
	var err error
	sr.levels = levels
	sr.diffs, err = calculateDiff(levels)

	if err != nil {
		return err
	}

	return nil
}

func (sr *SafetyReport) getSafetyReportSafety() bool {
	return sr.safe
}

func (sr *SafetyReport) setSafetyReportSafety(safety bool) {
	sr.safe = safety
}

func problemDampener(levels []int) bool {
	var results []bool
	var result bool
	var possibleSlices []int
	var diffs []int

	for i := 0; i < len(levels); i++ {
		possibleSlices = nil
		results = nil
		possibleSlices = append(possibleSlices, levels[:i]...)
		possibleSlices = append(possibleSlices, levels[i+1:]...)
		diffs, _ = calculateDiff(possibleSlices)

		result = checkForContinuity(diffs)
		results = append(results, result)

		result = checkForStaleness(diffs)
		results = append(results, result)

		result = checkForRapidChange(diffs)
		results = append(results, result)

		if slices.Contains(results, false) {
			continue
		} else {
			return true
		}

	}
	return false
}

func (sr *SafetyReport) SafetyCheck() (bool, error) {
	var checks []bool
	var result bool

	result = checkForContinuity(sr.diffs)
	checks = append(checks, result)

	result = checkForRapidChange(sr.diffs)
	checks = append(checks, result)

	result = checkForStaleness(sr.diffs)
	checks = append(checks, result)

	if slices.Contains(checks, false) && sr.ProblemDampener {
		fixedByProblemDampener := problemDampener(sr.levels)
		if fixedByProblemDampener {
			sr.setSafetyReportSafety(true)
			return true, nil
		}
	} else if slices.Contains(checks, false) {
		return false, nil
	} else {
		sr.setSafetyReportSafety(true)
		return true, nil
	}
	return false, nil

}

func checkForRapidChange(diffs []int) bool {
	for _, diff := range diffs {
		if math.Abs(float64(diff)) >= float64(4) {
			return false
		}
	}

	return true
}

func checkForContinuity(diffs []int) bool {
	var direction string
	for i, diff := range diffs {
		if i != 0 {
			if direction == "ascending" && diff < 0 {
				return false
			}
			if direction == "descending" && diff > 0 {
				return false
			}
		}
		if diff > 0 {
			direction = "ascending"
		} else if diff < 0 {
			direction = "descending"
		} else {
			direction = "stale"
		}
	}
	return true
}

func checkForStaleness(diffs []int) bool {
	for _, diff := range diffs {
		if diff == 0 {
			return false
		}
	}
	return true
}

func readFromFile(filePath string) ([]SafetyReport, error) {
	var safetyReports []SafetyReport

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scan1 := bufio.NewScanner(file)
	for scan1.Scan() {
		sLevels := strings.Split(scan1.Text(), " ")
		iLevels := make([]int, len(sLevels))
		for i, s := range sLevels {
			iLevels[i], _ = strconv.Atoi(s)
		}

		safetyReport := SafetyReport{ProblemDampener: false}
		err = safetyReport.SetSafetyLevels(iLevels)
		if err != nil {
			return nil, err
		}
		safetyReports = append(safetyReports, safetyReport)

	}

	return safetyReports, nil

}

func Main() {
	safetyReports, err := readFromFile("twentyFour/redNosedReports/redNosedReports.txt")
	if err != nil {
		log.Fatal(err)
	}

	safeReports := 0
	for _, report := range safetyReports {
		ok, _ := report.SafetyCheck()
		if ok {
			safeReports++
		}
	}
	fmt.Println(safeReports)

	safeReports = 0
	for _, report := range safetyReports {
		report.ProblemDampener = true
		ok, _ := report.SafetyCheck()
		if ok {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}

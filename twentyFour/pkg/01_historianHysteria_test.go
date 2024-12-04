package historianHysteria

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)
import "context"
import "github.com/cucumber/godog"

func Test01_historianHysteria(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: Initialize01_historianHysteriaScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../features/01_historianHysteria.feature"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func Initialize01_historianHysteriaScenario(ctx *godog.ScenarioContext) {

	ctx.Step(`^there are (.*) and (.*) to compare$`, thereAreFirstListAndSecondListToCompare)
	ctx.Step(`^I compare the lists$`, iCompareTheLists)
	ctx.Step(`^the totalDistance should equal (.*)$`, theTotalDistanceShouldEqualTotalDistance)
}

func thereAreFirstListAndSecondListToCompare(ctx context.Context, firstList string, secondList string) (context.Context, error) {
	var lsCompare ListCompare

	sFirstList := strings.Split(firstList, " ")
	iFirstList := make([]int, len(sFirstList))
	for i, s := range sFirstList {
		iFirstList[i], _ = strconv.Atoi(s)
	}

	sSecondList := strings.Split(secondList, " ")
	iSecondList := make([]int, len(sSecondList))
	for i, s := range sSecondList {
		iSecondList[i], _ = strconv.Atoi(s)
	}
	lsCompare.FirstList = iFirstList
	lsCompare.SecondList = iSecondList

	return context.WithValue(ctx, "lsCompare", lsCompare), nil
}

func iCompareTheLists(ctx context.Context) (context.Context, error) {
	lsCompare := ctx.Value("lsCompare").(ListCompare)
	lsCompare.CalculateDistance()

	return context.WithValue(ctx, "totalDistance", lsCompare.totalDistance), nil
}
func theTotalDistanceShouldEqualTotalDistance(ctx context.Context, totalDistance int) error {
	if ctx.Value("totalDistance").(int) != totalDistance {
		return fmt.Errorf("expected distance to be %d but ther was %d", totalDistance, 1)
	}

	return nil
}

package historianHysteria

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)
import "context"
import "github.com/cucumber/godog"

func TestHistorianHysteria(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeHistorianHysteriaScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"historianHysteria.feature"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeHistorianHysteriaScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^there are (.*) and (.*) to compare$`, thereAreFirstListAndSecondListToCompare)
	ctx.Step(`^I compare the lists for total distance$`, iCompareTheListsForTotalDistance)
	ctx.Step(`^the totalDistance should equal (.*)$`, theTotalDistanceShouldEqualTotalDistance)
	ctx.Step(`^I compare the lists for similarity$`, iCompareTheListsForSimilarity)
	ctx.Step(`^the similarity should equal (.*)$`, theSimilarityShouldEqualSimilarity)
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

func iCompareTheListsForTotalDistance(ctx context.Context) (context.Context, error) {
	lsCompare := ctx.Value("lsCompare").(ListCompare)
	lsCompare.CalculateDistance()

	return context.WithValue(ctx, "totalDistance", lsCompare.TotalDistance), nil
}

func theTotalDistanceShouldEqualTotalDistance(ctx context.Context, totalDistance int) error {
	if ctx.Value("totalDistance").(int) != totalDistance {
		return fmt.Errorf("expected distance to be %d but ther was %d", totalDistance, 1)
	}

	return nil
}
func iCompareTheListsForSimilarity(ctx context.Context) (context.Context, error) {
	lsCompare := ctx.Value("lsCompare").(ListCompare)
	lsCompare.CalculateSimilarity()

	return context.WithValue(ctx, "totalSimilarity", lsCompare.TotalSimilarity), nil
}
func theSimilarityShouldEqualSimilarity(ctx context.Context, similarity int) error {
	if ctx.Value("totalSimilarity").(int) != similarity {
		return fmt.Errorf("expected similarity to be %d but ther was %d", similarity, 1)
	}
	return nil
}

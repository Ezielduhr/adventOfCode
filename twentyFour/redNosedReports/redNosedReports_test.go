package redNosedReports

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)
import "context"
import "github.com/cucumber/godog"

func TestRedNosedReports(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeRedNosedReportsScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"redNosedReports.feature"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeRedNosedReportsScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^there are safety (.*)$`, thereAreSafetyReports)
	ctx.Step(`^we enable the problem dampener$`, weEnableTheProblemDampener)
	ctx.Step(`^we check the report$`, weCheckTheReport)
	ctx.Step(`^the report should be (.*)$`, theReportShouldBeResult)
}

func thereAreSafetyReports(ctx context.Context, levels string) (context.Context, error) {
	sLevels := strings.Split(levels, " ")
	iLevels := make([]int, len(sLevels))
	for i, s := range sLevels {
		iLevels[i], _ = strconv.Atoi(s)
	}

	safetyReport := SafetyReport{ProblemDampener: false}
	err := safetyReport.SetSafetyLevels(iLevels)
	if err != nil {
		return ctx, err
	}
	return context.WithValue(ctx, "safetyReport", safetyReport), nil
}

func weEnableTheProblemDampener(ctx context.Context) (context.Context, error) {
	safetyReport := ctx.Value("safetyReport").(SafetyReport)
	safetyReport.ProblemDampener = true

	return context.WithValue(ctx, "safetyReport", safetyReport), nil
}

func weCheckTheReport(ctx context.Context) (context.Context, error) {
	safetyReport := ctx.Value("safetyReport").(SafetyReport)
	if _, err := safetyReport.SafetyCheck(); err != nil {
		return ctx, fmt.Errorf("safety check failed: %w", err)
	}

	return context.WithValue(ctx, "safetyReport", safetyReport), nil
}

func theReportShouldBeResult(ctx context.Context, result string) (context.Context, error) {
	bResult := false
	if result == "safe" {
		bResult = true
	}

	safetyReport := ctx.Value("safetyReport").(SafetyReport)
	isSafe := safetyReport.getSafetyReportSafety()

	if safetyReport.getSafetyReportSafety() != bResult {
		return ctx, fmt.Errorf("safety check failed: expected safe to be %t, got %t", bResult, isSafe)
	} else {
		return ctx, nil
	}
}

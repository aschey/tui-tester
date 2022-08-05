package test

import (
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	tuitest "github.com/aschey/tui-tester"
	"github.com/stretchr/testify/suite"
)

func init() {
	os.Args = append(os.Args, "-tuicover", "-tuicoverpkg", "../...")
}

type TestSuite struct {
	suite.Suite
	suite   *tuitest.Suite
	console *tuitest.Console
}

func cleanupCoverageFile() {
	if _, err := os.Stat("coverage.out"); err == nil {
		os.Remove("coverage.out")
	} else if !errors.Is(err, os.ErrNotExist) {
		panic("Error reading coverage file " + err.Error())
	}
}

func (suite *TestSuite) SetupSuite() {
	cleanupCoverageFile()
}

func (suite *TestSuite) setup(opts ...tuitest.Option) {
	testSuite := tuitest.NewSuite()
	tester, err := testSuite.NewTester("./testapp", opts...)
	suite.suite = testSuite

	suite.Require().NoError(err)
	console, err := tester.CreateConsole()
	suite.Require().NoError(err)
	suite.console = console
	console.TrimOutput = true

	_, err = suite.console.WaitFor(func(state tuitest.TermState) bool {
		return state.Output() == "You typed:"
	})
	suite.Require().NoError(err)
}

func (suite *TestSuite) TearDownTest() {
	suite.console.SendString(tuitest.KeyCtrlC)
	suite.Require().NoError(suite.console.WaitForTermination())
	suite.Require().NoError(suite.suite.TearDown())
	fileBytes, err := os.ReadFile("coverage.out")
	suite.Require().NoError(err)
	fileStr := string(fileBytes)
	suite.Require().True(strings.HasPrefix(fileStr, "mode: atomic"))
}

func (suite *TestSuite) TearDownSuite() {
	cleanupCoverageFile()
}

func TestTuiTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestSendInputAndExpectOutput() {
	suite.setup(tuitest.WithErrorHandler(func(err error) error {
		suite.Require().NoError(err)
		return err
	}))

	suite.console.SendString("input")
	_, _ = suite.console.WaitFor(func(state tuitest.TermState) bool {
		return state.Output() == "You typed: input"
	})
}

func (suite *TestSuite) TestMinInputInterval() {
	suite.setup(tuitest.WithMinInputInterval(100 * time.Millisecond))

	start := time.Now()
	suite.console.SendString("a")
	suite.console.SendString("b")
	end := time.Now()
	duration := end.Sub(start)
	suite.Require().True(duration >= 100*time.Millisecond)
	suite.Require().True(duration < 200*time.Millisecond)
}

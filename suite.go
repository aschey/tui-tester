package tuitest

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"

	"golang.org/x/tools/cover"
)

var collectCoverage *bool
var coverpkg *string
var coverageFile *string

func init() {
	collectCoverage = flag.Bool("tuicover", false, "")
	coverpkg = flag.String("tuicoverpkg", ".", "")
	coverageFile = flag.String("tuicoverfile", "coverage.out", "")
}

type Suite struct {
	testers         []*Tester
	coverageFile    string
	collectCoverage bool
}

func NewSuite() *Suite {
	return &Suite{testers: []*Tester{}, coverageFile: *coverageFile, collectCoverage: *collectCoverage}
}

func (s *Suite) NewTester(binDir string, opts ...Option) (*Tester, error) {
	tmpdir, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, err
	}
	exe_name := "test_bin"
	exePath := path.Join(tmpdir, exe_name)

	var buildTestCmd *exec.Cmd
	if *collectCoverage {
		buildTestCmd = exec.Command("go", "test", binDir, "-covermode=atomic", "-c", "-o", exePath, "-coverpkg", *coverpkg)
	} else {
		buildTestCmd = exec.Command("go", "test", binDir, "-c", "-o", exePath)
	}

	output, err := buildTestCmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf(string(output))
	}

	tester := &Tester{
		exePath:            exePath,
		collectCoverage:    *collectCoverage,
		defaultWaitTimeout: time.Second,
		minInputInterval:   time.Millisecond,
		terminationTimeout: time.Second,
		onError:            func(err error) error { return err },
	}

	for _, opt := range opts {
		if err := opt(tester); err != nil {
			return nil, err
		}
	}
	s.testers = append(s.testers, tester)
	return tester, nil
}

func (s *Suite) TearDown() error {
	coverageFiles := []*cover.Profile{}
	for _, tester := range s.testers {
		if s.collectCoverage {
			files, err := tester.getCoverageFiles()
			if err != nil {
				return err
			}
			coverageFiles = append(coverageFiles, files...)
		}
		err := tester.tearDown()
		if err != nil {
			return err
		}
	}
	covFile, err := os.Create(s.coverageFile)
	if err != nil {
		return err
	}
	err = dumpProfiles(coverageFiles, covFile)
	if err != nil {
		return err
	}
	return covFile.Close()
}

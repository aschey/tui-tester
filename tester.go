package tuitest

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/aschey/termtest"
	"github.com/aschey/termtest/expect"
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

type Tester struct {
	exePath         string
	coverageFile    string
	collectCoverage bool
}

func (t Tester) NewConsole(args []string) (*Console, error) {
	defaultTimeout := time.Second

	binDir := path.Dir(t.exePath)
	if t.collectCoverage {
		tempFile, err := os.CreateTemp(binDir, "*.cov")
		if err != nil {
			return nil, err
		}

		args = append(args, "-test.coverprofile")
		args = append(args, tempFile.Name())
	}

	opts := termtest.Options{
		CmdName: t.exePath,
		Args:    args,
		ExtraOpts: []expect.ConsoleOpt{
			expect.WithDefaultExpectTimeout(5 * time.Second),
		},
	}
	console, err := termtest.New(opts)
	if err != nil {
		return nil, err
	}

	tester := Console{
		console:   console,
		lastInput: time.Now(),
		Timeout:   defaultTimeout,
		last:      "",
	}

	return &tester, nil
}

func NewTester(binDir string) (Tester, error) {
	tmpdir, err := os.MkdirTemp("", "")
	if err != nil {
		return Tester{}, err
	}
	exePath := path.Join(tmpdir, "instr_bin")

	var buildTestCmd *exec.Cmd
	if *collectCoverage {
		buildTestCmd = exec.Command("go", "test", binDir, "-covermode=atomic", "-c", "-o", exePath, "-coverpkg", *coverpkg)
	} else {
		buildTestCmd = exec.Command("go", "test", binDir, "-c", "-o", exePath)
	}

	output, err := buildTestCmd.CombinedOutput()
	if err != nil {
		return Tester{}, fmt.Errorf(string(output))
	}

	return Tester{exePath: exePath, collectCoverage: *collectCoverage, coverageFile: *coverageFile}, nil
}

func (t Tester) TearDown() error {
	binDir := path.Dir(t.exePath)
	if t.collectCoverage {
		files, err := os.ReadDir(binDir)
		if err != nil {
			return err
		}
		merged := []*cover.Profile{}
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".cov") {
				profiles, err := cover.ParseProfiles(path.Join(binDir, f.Name()))
				if err != nil {
					return err
				}
				for _, p := range profiles {
					merged = addProfile(merged, p)
				}
			}
		}
		covFile, err := os.Create(t.coverageFile)
		if err != nil {
			return err
		}
		dumpProfiles(merged, covFile)
	}

	err := os.RemoveAll(binDir)
	return err
}

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
	exePath            string
	coverageFile       string
	collectCoverage    bool
	defaultWaitTimeout time.Duration
	minInputInterval   time.Duration
	terminationTimeout time.Duration
	onError            func(err error) error
}

func (t *Tester) CreateConsole(args []string) (*Console, error) {
	binDir := path.Dir(t.exePath)
	if t.collectCoverage {
		tempFile, err := os.CreateTemp(binDir, "*.cov")
		if err != nil {
			return nil, err
		}

		args = append(args, "-test.coverprofile", tempFile.Name())
	}

	opts := termtest.Options{
		CmdName: t.exePath,
		Args:    args,
		ExtraOpts: []expect.ConsoleOpt{
			expect.WithDefaultExpectTimeout(t.defaultWaitTimeout),
		},
	}
	consoleProcess, err := termtest.New(opts)
	if err != nil {
		return nil, err
	}

	console := &Console{
		consoleProcess:     consoleProcess,
		lastInput:          time.Now(),
		minInputInterval:   t.minInputInterval,
		terminationTimeout: t.terminationTimeout,
		onError:            t.onError,
		last:               "",
	}

	return console, nil
}

func NewTester(binDir string, opts ...Option) (*Tester, error) {
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
		coverageFile:       *coverageFile,
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
	return tester, nil
}

func (t *Tester) TearDown() error {
	binDir := path.Dir(t.exePath)
	if t.collectCoverage {
		files, err := os.ReadDir(binDir)
		if err != nil {
			return t.onError(err)
		}
		merged := []*cover.Profile{}
		for _, f := range files {
			if strings.HasSuffix(f.Name(), ".cov") {
				profiles, err := cover.ParseProfiles(path.Join(binDir, f.Name()))
				if err != nil {
					return t.onError(err)
				}
				for _, p := range profiles {
					merged = addProfile(merged, p)
				}
			}
		}
		covFile, err := os.Create(t.coverageFile)
		if err != nil {
			return t.onError(err)
		}
		dumpProfiles(merged, covFile)
	}

	err := os.RemoveAll(binDir)
	return t.onError(err)
}

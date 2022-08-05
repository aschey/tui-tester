package tuitest

import (
	"io/fs"
	"os"
	"path"
	"strings"
	"time"

	"github.com/aschey/termtest"
	"github.com/aschey/termtest/expect"
	"golang.org/x/tools/cover"
)

type Tester struct {
	exePath            string
	collectCoverage    bool
	defaultWaitTimeout time.Duration
	minInputInterval   time.Duration
	terminationTimeout time.Duration
	onError            func(err error) error
}

func (t *Tester) CreateConsole(args ...string) (*Console, error) {
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
		minInputInterval:   t.minInputInterval,
		terminationTimeout: t.terminationTimeout,
		onError:            t.onError,
		last:               "",
	}

	return console, nil
}

func (t *Tester) getCoverageFiles() ([]*cover.Profile, error) {
	binDir := path.Dir(t.exePath)
	if t.collectCoverage {
		files, err := os.ReadDir(binDir)
		if err != nil {
			return nil, t.onError(err)

		}
		merged, err := t.profiles(files, binDir)
		if err != nil {
			return nil, t.onError(err)
		}
		return merged, nil
	}
	return nil, nil
}

func (t *Tester) tearDown() error {
	binDir := path.Dir(t.exePath)

	err := os.RemoveAll(binDir)
	if err != nil {
		return t.onError(err)
	}
	return nil
}

func (t *Tester) profiles(files []fs.DirEntry, binDir string) ([]*cover.Profile, error) {
	merged := []*cover.Profile{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".cov") {
			profiles, err := cover.ParseProfiles(path.Join(binDir, f.Name()))
			if err != nil {
				return nil, err
			}
			for _, p := range profiles {
				merged = addProfile(merged, p)
			}
		}
	}
	return merged, nil
}

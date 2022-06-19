package tuitest

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/aschey/termtest"
	"golang.org/x/tools/cover"
)

type Tester struct {
	exePath      string
	coverageFile string
}

func (t Tester) NewConsole(args []string) (*Console, error) {
	defaultTimeout := time.Second

	binDir := path.Dir(t.exePath)
	tempFile, err := os.CreateTemp(binDir, "*.cov")
	if err != nil {
		return nil, err
	}
	println(tempFile.Name())
	args = append(args, "-test.coverprofile")
	args = append(args, tempFile.Name())
	opts := termtest.Options{
		CmdName: t.exePath,
		Args:    args,
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

func NewTester(binDir string, coverPkg string, coverageFile string) (Tester, error) {
	tmpdir, err := os.MkdirTemp("", "")
	if err != nil {
		return Tester{}, err
	}
	exePath := path.Join(tmpdir, "instr_bin")

	buildTestCmd := exec.Command("go", "test", binDir, "-covermode=atomic", "-c", "-o", exePath, "-coverpkg", coverPkg+"/...")
	output, err := buildTestCmd.CombinedOutput()
	if err != nil {
		return Tester{}, fmt.Errorf(string(output))
	}

	return Tester{exePath: exePath, coverageFile: coverageFile}, nil
}

func (t Tester) CoverageTearDown() error {
	binDir := path.Dir(t.exePath)
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
	err = os.RemoveAll(binDir)
	return err
}

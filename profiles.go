package tuitest

import (
	"fmt"
	"os"
	"sort"

	"golang.org/x/tools/cover"
)

func addProfile(profiles []*cover.Profile, p *cover.Profile) []*cover.Profile {
	i := sort.Search(len(profiles), func(i int) bool { return profiles[i].FileName >= p.FileName })
	profiles = append(profiles, nil)
	copy(profiles[i+1:], profiles[i:])
	profiles[i] = p
	return profiles
}

func dumpProfiles(profiles []*cover.Profile, out *os.File) error {
	if len(profiles) == 0 {
		return nil
	}
	stat, err := out.Stat()
	if err != nil {
		return err
	}
	if stat.Size() == 0 {
		fmt.Fprintf(out, "mode: %s\n", profiles[0].Mode)
	}

	for _, p := range profiles {
		for _, b := range p.Blocks {
			fmt.Fprintf(out, "%s:%d.%d,%d.%d %d %d\n", p.FileName, b.StartLine, b.StartCol, b.EndLine, b.EndCol, b.NumStmt, b.Count)
		}
	}
	return nil
}

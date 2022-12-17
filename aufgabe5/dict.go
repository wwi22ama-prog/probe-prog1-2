package aufgabe5

import (
	"fmt"
	"strings"
)

type Entry struct {
	De string
	En string
}

func (entry Entry) String() string {
	return fmt.Sprintf("%s : %s", entry.De, entry.En)
}

type Dictionary struct {
	Entries []Entry
}

func (dict Dictionary) String() string {
	lines := []string{}
	for i, entry := range dict.Entries {
		lines = append(lines, fmt.Sprintf("%d - %s", i, entry))
	}
	return strings.Join(lines, "\n")
}

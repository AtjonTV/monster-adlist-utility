package monster

import "fmt"

type Monster struct {
	Sources    Sources
	VerboseLog bool
}

func New(sources Sources, printVerbose bool) Monster {
	return Monster{
		Sources:    sources,
		VerboseLog: printVerbose,
	}
}

func (m *Monster) DebugLog(format string, a ...any) {
	if m.VerboseLog {
		fmt.Printf(format, a...)
	}
}

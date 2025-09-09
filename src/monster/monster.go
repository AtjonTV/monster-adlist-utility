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

func (m *Monster) SetRewriteFlag(forceEnable bool, forceDisable bool) {
	m.Sources.Rewrite.Enable = forceEnable || (m.Sources.Rewrite.Enable && !forceDisable)
}

func (m *Monster) SetCleanFlag(forceEnable bool, forceDisable bool) {
	m.Sources.CleanRule.Enable = forceEnable || (m.Sources.CleanRule.Enable && !forceDisable)
}

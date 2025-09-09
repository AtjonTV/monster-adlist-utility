/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Monster struct {
	Sources    Sources
	VerboseLog bool
	OutputDir  string
}

func New(sources Sources, printVerbose bool) Monster {
	return Monster{
		Sources:    sources,
		VerboseLog: printVerbose,
	}
}

func NewFromFile(path string, printVerbose bool) (Monster, error) {
	var data, err = os.ReadFile(path)
	if err != nil {
		return Monster{}, err
	}

	var sources Sources
	err = yaml.Unmarshal(data, &sources)
	if err != nil {
		return Monster{}, err
	}

	return New(sources, printVerbose), nil
}

func (m *Monster) SetOutputDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return err
	}
	m.OutputDir = dir
	return nil
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

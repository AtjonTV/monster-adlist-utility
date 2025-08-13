/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import "fmt"

type Sources struct {
	Rewrite   Rewrite      `yaml:"rewrite"`
	CleanRule CleanRule    `yaml:"cleanup"`
	Allow     []SourceList `yaml:"allow"`
	Block     []SourceList `yaml:"block"`

	VerboseLog bool
}

func (sources *Sources) DebugLog(format string, a ...any) {
	if sources.VerboseLog {
		fmt.Printf(format, a...)
	}
}

type Rewrite struct {
	Enable   bool   `yaml:"enable"`
	CustomIP string `yaml:"custom_ip"`
	Mode     string `yaml:"mode"`
}

type CleanRule struct {
	Enable   bool `yaml:"enable"`
	KeepDays int  `yaml:"keep_days"` // only keep N days (0 = disabled, >=1 = N days into the past)
}

type SourceList struct {
	Name   string     `yaml:"name"`
	Url    string     `yaml:"url"`
	Type   string     `yaml:"type"`
	Data   []string   `yaml:"data"`
	Header ListHeader `yaml:"header"`
	Trim   TrimInfo   `yaml:"trim"`

	TempFile string
}

type ListHeader struct {
	Title    string `yaml:"title"`
	Homepage string `yaml:"homepage"`
	License  string `yaml:"license"`
}

type TrimInfo struct {
	Head int `yaml:"head"`
	Tail int `yaml:"tail"`
}

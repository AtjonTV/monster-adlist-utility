/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

type Sources struct {
	Allow []SourceList `yaml:"allow"`
	Block []SourceList `yaml:"block"`
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

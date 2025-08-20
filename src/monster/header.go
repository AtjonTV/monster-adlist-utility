/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import "fmt"

func (m *Monster) RenderHeader(entryCount int) []string {
	var header = make([]string, 0, len(m.Sources.Allow)+len(m.Sources.Block)+4)

	header = append(header, "#")
	header = append(header, "# Title: AtjonTV's Monster Adlist")
	header = append(header, "# Author: Thomas Obernosterer")
	header = append(header, "# Homepage: https://monster-adlist.atvg.cloud/")
	header = append(header, fmt.Sprintf("# Total of %d unique domains from %d lists", entryCount, len(m.Sources.Allow)+len(m.Sources.Block)))
	header = append(header, "# Based on:")

	var allSources = make([]SourceList, 0, len(m.Sources.Allow)+len(m.Sources.Block))
	allSources = append(allSources, m.Sources.Allow...)
	allSources = append(allSources, m.Sources.Block...)

	for _, source := range allSources {
		if source.Header.Title == "" {
			continue
		}
		header = append(header, "#  Title: "+source.Header.Title+"")
		if source.Header.Homepage != "" {
			header = append(header, "#    Homepage: "+source.Header.Homepage+"")
		}
		if source.Header.License != "" {
			header = append(header, "#    License: "+source.Header.License+"")
		}
	}
	header = append(header, "#")

	return header
}

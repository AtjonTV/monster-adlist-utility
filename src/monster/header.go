package monster

func RenderHeader(sources *Sources) []string {
	var header = make([]string, 0, len(sources.Allow)+len(sources.Block)+4)

	header = append(header, "#")
	header = append(header, "# Title: AtjonTV's Monster Adlist")
	header = append(header, "# Author: Thomas Obernosterer")
	header = append(header, "# Homepage: https://monster-adlist.atvg.cloud/")
	header = append(header, "# Based on:")

	var allSources = make([]SourceList, 0, len(sources.Allow)+len(sources.Block))
	allSources = append(allSources, sources.Allow...)
	allSources = append(allSources, sources.Block...)

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

/*
 * Copyright (c) 2025 Thomas Obernosterer. Licensed under the EUPL-1.2 or later.
 *
 * SPDX-License-Identifier: EUPL-1.2-or-later
 */

package monster

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (m *Monster) CleanUp() error {
	if !m.Sources.CleanRule.Enable {
		return nil
	}

	if m.Sources.CleanRule.KeepDays > 0 {
		fmt.Printf("CLEAN: preparing clean-up of files older %d days in '%s'\n", m.Sources.CleanRule.KeepDays, m.OutputDir)
		var files, err = os.ReadDir(m.OutputDir)
		if err != nil {
			return err
		}
		var removalList = make([]string, 0, len(files))
		for _, file := range files {
			fileTime, err := extractTimeFromFileName(file.Name())
			if err != nil {
				m.DebugLog("DEBUG: could not extract time from file '%s'\n", file.Name())
				continue
			}
			var hrs = time.Since(fileTime).Round(24 * time.Hour).Hours()
			var oldHrs = float64(m.Sources.CleanRule.KeepDays) * 24
			m.DebugLog("DEBUG: file='%s' time=%s; hours=%f; considered old=%f\n", file.Name(), fileTime.Format(time.RFC3339), hrs, oldHrs)
			if hrs >= oldHrs {
				removalList = append(removalList, file.Name())
			}
		}
		fmt.Printf("CLEAN: removing %d files\n", len(removalList))
		for _, file := range removalList {
			fileName := m.OutputDir + pathSeparator + file
			m.DebugLog("DEBUG: removing file '%s'\n", fileName)
			_ = os.Remove(fileName)
		}
		fmt.Printf("CLEAN: done\n")
	}

	return nil
}

func extractTimeFromFileName(fileName string) (time.Time, error) {
	cleanString := strings.TrimPrefix(fileName, "monster_")
	suffixes := []string{".list", ".update", "_rewrite.list"}
	for _, suffix := range suffixes {
		if cut, hadSuffix := strings.CutSuffix(cleanString, suffix); hadSuffix {
			cleanString = cut
			break
		}
	}

	return time.Parse("2006-01-02_15-04", cleanString)
}

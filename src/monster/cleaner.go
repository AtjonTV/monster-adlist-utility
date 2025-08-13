/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func CleanUp(sources *Sources, outDir string) error {
	if sources.CleanRule.KeepDays > 0 {
		fmt.Printf("CLEAN: preparing clean-up of files older %d days in '%s'\n", sources.CleanRule.KeepDays, outDir)
		var files, err = os.ReadDir(outDir)
		if err != nil {
			return err
		}
		var removalList = make([]string, 0, len(files))
		for _, file := range files {
			fileTime, err := extractTimeFromFileName(file.Name())
			if err != nil {
				continue
			}
			var hrs = time.Since(fileTime).Round(24 * time.Hour).Hours()
			var oldHrs = float64(sources.CleanRule.KeepDays) * 24
			sources.DebugLog("DEBUG: time=%s; hours=%f; considered old=%f\n", fileTime.Format(time.RFC3339), hrs, oldHrs)
			if hrs >= oldHrs {
				removalList = append(removalList, file.Name())
			}
		}
		fmt.Printf("CLEAN: removing %d files\n", len(removalList))
		for _, file := range removalList {
			fileName := outDir + pathSeparator + file
			sources.DebugLog("DEBUG: removing file '%s'\n", fileName)
			_ = os.Remove(fileName)
		}
		fmt.Printf("CLEAN: done\n")
	}

	return nil
}

func extractTimeFromFileName(fileName string) (time.Time, error) {
	cleanString := strings.TrimPrefix(fileName, "monster_")
	var suffix = ".list"
	if strings.HasSuffix(cleanString, ".update") {
		suffix = ".update"
	} else if strings.HasSuffix(cleanString, "_rewrite.list") {
		suffix = "_rewrite.list"
	}

	cleanName := strings.TrimSuffix(cleanString, suffix)
	return time.Parse("2006-01-02_15-04", cleanName)
}

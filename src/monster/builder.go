/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func BuildMonster(sources *Sources, outFolder string) (string, error) {
	blockList, err := buildSubList(sources.Block)
	if err != nil {
		return "", err
	}

	allowList, err := buildSubList(sources.Allow)
	if err != nil {
		return "", err
	}

	var monsterList = RenderHeader(sources)
	monsterList = append(monsterList, removeAllowFromBlock(blockList, allowList)...)

	var now = time.Now().Format("2006-01-02_15-04")
	var monsterName = fmt.Sprintf("%s%smonster_%s.list", outFolder, pathSeparator, now)
	err = writeFile(monsterName, []byte(strings.Join(monsterList, "\n")))
	if err != nil {
		return "", err
	}

	fmt.Printf("MONSTER: created monster file: %s\n", monsterName)
	return monsterName, nil
}

func removeAllowFromBlock(blockList []string, allowList []string) []string {
	elementsToRemove := make(map[string]bool)
	for _, element := range allowList {
		elementsToRemove[element] = true
	}

	result := make([]string, 0, len(blockList))

	for _, element := range blockList {
		if !elementsToRemove[element] && element != "" {
			result = append(result, element)
		}
	}
	return result
}

func buildSubList(lists []SourceList) ([]string, error) {
	var files []string
	for _, list := range lists {
		files = append(files, list.TempFile)
	}
	defer func() {
		for _, file := range files {
			_ = os.Remove(file)
		}
	}()

	var allLines []string
	for _, file := range files {
		fileHandle, err := os.Open(file)
		if err != nil {
			return []string{}, err
		}

		scanner := bufio.NewScanner(fileHandle)
		for scanner.Scan() {
			allLines = append(allLines, scanner.Text())
		}

		err = fileHandle.Close()
		if err != nil {
			return []string{}, err
		}
	}

	sort.Strings(allLines)

	sublist := removeDuplicates(allLines)
	return sublist, nil
}

func removeDuplicates(elements []string) []string {
	uniqueLines := make([]string, 0, len(elements))
	if len(elements) > 0 {
		uniqueLines = append(uniqueLines, elements[0]) // add the first line
	}
	for i := 1; i < len(elements); i++ {
		if elements[i] != elements[i-1] {
			uniqueLines = append(uniqueLines, elements[i])
		}
	}

	return uniqueLines
}

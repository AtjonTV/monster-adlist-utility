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

	var domainList = removeAllowFromBlock(blockList, allowList)
	var monsterList = make([]string, 0, 30+len(domainList))

	var writeNormalMonster = true

	if sources.Rewrite.Enable {
		var rewriteList = make([]string, 0, 30+len(domainList))
		rewriteList = append(rewriteList, RenderHeader(sources)...)
		for _, domain := range domainList {
			rewriteList = append(rewriteList, sources.Rewrite.CustomIP+" "+domain)
		}
		switch sources.Rewrite.Mode {
		case "new_file":
			monsterName, err := writeListToFile(outFolder, "_rewrite", rewriteList)
			if err != nil {
				return "", err
			}
			fmt.Printf("MONSTER: created monster rewrite file: %s\n", monsterName)
			break
		case "override":
			writeNormalMonster = false
			monsterList = rewriteList
			break
		}
		rewriteList = nil
	}

	if writeNormalMonster {
		monsterList = append(monsterList, RenderHeader(sources)...)
		monsterList = append(monsterList, domainList...)
	}
	domainList = nil

	monsterName, err := writeListToFile(outFolder, "", monsterList)
	if err != nil {
		return "", err
	}

	fmt.Printf("MONSTER: created monster file: %s\n", monsterName)
	return monsterName, nil
}

func writeListToFile(folder string, suffix string, list []string) (string, error) {
	var now = time.Now().Format("2006-01-02_15-04")
	var monsterName = fmt.Sprintf("%s%smonster_%s%s.list", folder, pathSeparator, now, suffix)
	err := writeFile(monsterName, []byte(strings.Join(list, "\n")))
	if err != nil {
		return "", err
	}
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

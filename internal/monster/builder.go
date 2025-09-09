/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"atjon.tv/monster/internal/utils"
)

func (m *Monster) BuildMonsterList() (string, error) {
	blockList, err := buildSubList(m.Sources.Block)
	if err != nil {
		return "", err
	}

	allowList, err := buildSubList(m.Sources.Allow)
	if err != nil {
		return "", err
	}

	var domainList = utils.RemoveListItems(blockList, allowList)
	blockList = nil // needs to be freed here, currently has 90+MB (2025-09-09)
	allowList = nil
	var monsterList = make([]string, 0, 30+len(domainList))

	var writeNormalMonster = true

	if m.Sources.Rewrite.Enable {
		var rewriteList = make([]string, 0, 30+len(domainList))
		rewriteList = append(rewriteList, m.RenderHeader(len(domainList))...)
		for _, domain := range domainList {
			rewriteList = append(rewriteList, m.Sources.Rewrite.CustomIP+" "+domain)
		}
		switch m.Sources.Rewrite.Mode {
		case "new_file":
			monsterName, err := writeListToFile(m.OutputDir, "_rewrite", rewriteList)
			if err != nil {
				return "", err
			}
			fmt.Printf("MONSTER: created monster rewrite file: %s\n", monsterName)
		case "override":
			writeNormalMonster = false
			monsterList = rewriteList
		}
		rewriteList = nil
	}

	if writeNormalMonster {
		monsterList = append(monsterList, m.RenderHeader(len(domainList))...)
		monsterList = append(monsterList, domainList...)
	}
	domainList = nil

	monsterName, err := writeListToFile(m.OutputDir, "", monsterList)
	if err != nil {
		return "", err
	}

	fmt.Printf("MONSTER: created monster file: %s\n", monsterName)
	return monsterName, nil
}

func writeListToFile(folder string, suffix string, list []string) (string, error) {
	var now = time.Now().Format("2006-01-02_15-04")
	var monsterName = fmt.Sprintf("%s%smonster_%s%s.list", folder, pathSeparator, now, suffix)
	err := utils.WriteDataToFile(monsterName, []byte(strings.Join(list, "\n")))
	if err != nil {
		return "", err
	}
	return monsterName, nil
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
		lines, err := utils.ReadLinesFromFile(file)
		if err != nil {
			return []string{}, err
		}
		allLines = append(allLines, lines...)
	}

	sort.Strings(allLines)

	sublist := utils.RemoveDuplicatesFromList(allLines)
	return sublist, nil
}

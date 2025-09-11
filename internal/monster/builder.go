/*
 * Copyright (c) 2025 Thomas Obernosterer. Licensed under the EUPL-1.2 or later.
 *
 * SPDX-License-Identifier: EUPL-1.2-or-later
 */

package monster

import (
	"fmt"
	"os"
	"runtime"
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

	utils.RemoveListItemsMut(&blockList, &allowList)
	var monsterList = make([]string, 0, 30+len(blockList))

	var writeNormalMonster = true

	if m.Sources.Rewrite.Enable {
		var rewriteList = make([]string, 0, 30+len(blockList))
		rewriteList = append(rewriteList, m.RenderHeader(len(blockList))...)
		for _, domain := range blockList {
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
		monsterList = append(monsterList, m.RenderHeader(len(blockList))...)
		monsterList = append(monsterList, blockList...)
	}

	monsterName, err := writeListToFile(m.OutputDir, "", monsterList)
	if err != nil {
		return "", err
	}

	// set to nil, so it can be garbage collected
	blockList = nil
	allowList = nil
	monsterList = nil
	// run garbage collector
	runtime.GC()

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

	allLines = nil
	runtime.GC()

	return sublist, nil
}

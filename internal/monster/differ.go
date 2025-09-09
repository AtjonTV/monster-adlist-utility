/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import (
	"fmt"
	"os"
	"strings"

	"atjon.tv/monster/internal/utils"
)

func (m *Monster) CreateDiffFile(previousMonster string, newMonster string) error {
	_, err := os.Stat(previousMonster)
	if os.IsNotExist(err) {
		return fmt.Errorf("the previous monster file was not found: %s", previousMonster)
	}

	fmt.Printf("DIFF: preparing diff file creation between '%s' and '%s'\n", previousMonster, newMonster)

	prevList, err := utils.ReadLinesFromFile(previousMonster)
	if err != nil {
		return err
	}
	newList, err := utils.ReadLinesFromFile(newMonster)
	if err != nil {
		return err
	}

	trimLines(prevList)
	removeComments(prevList)
	trimLines(newList)
	removeComments(newList)

	domainList := utils.RemoveListItems(newList, prevList)
	newList = nil
	prevList = nil

	var diffList = m.RenderHeader(len(domainList))
	var headerSize = len(diffList)
	diffList = append(diffList, domainList...)
	domainList = nil

	if len(diffList) == headerSize {
		fmt.Println("DIFF: no changes detected, skipping file creation.")
		return nil
	}

	var patchName = strings.ReplaceAll(newMonster, ".list", ".update")
	err = utils.WriteDataToFile(patchName, []byte(strings.Join(diffList, "\n")))
	if err != nil {
		return err
	}

	fmt.Printf("DIFF: created patch file: %s\n", patchName)
	return nil
}

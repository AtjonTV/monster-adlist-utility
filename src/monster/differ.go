/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import (
	"fmt"
	"os"
	"strings"
)

func CreateDiffFile(sources *Sources, previousMonster string, newMonster string) error {
	_, err := os.Stat(previousMonster)
	if os.IsNotExist(err) {
		return fmt.Errorf("the previous monster file was not found: %s", previousMonster)
	}

	fmt.Printf("DIFF: preparing diff file creation between '%s' and '%s'\n", previousMonster, newMonster)

	prevList, err := readListData(previousMonster)
	if err != nil {
		return err
	}
	newList, err := readListData(newMonster)
	if err != nil {
		return err
	}

	trimLines(prevList)
	removeComments(prevList)
	trimLines(newList)
	removeComments(newList)

	domainList := removeAllowFromBlock(newList, prevList)

	var diffList = RenderHeader(sources, len(domainList))
	var headerSize = len(diffList)
	diffList = append(diffList, domainList...)
	newList = nil
	prevList = nil
	domainList = nil

	if len(diffList) == headerSize {
		fmt.Println("DIFF: no changes detected, skipping file creation.")
		return nil
	}

	var patchName = strings.Replace(newMonster, ".list", ".update", -1)
	err = writeFile(patchName, []byte(strings.Join(diffList, "\n")))
	if err != nil {
		return err
	}

	fmt.Printf("DIFF: created patch file: %s\n", patchName)
	return nil
}

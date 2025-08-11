/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import (
	"fmt"
	"os"
	"strings"
)

func CreatePatch(sources *Sources, previousMonster string, newMonster string) error {
	_, err := os.Stat(previousMonster)
	if os.IsNotExist(err) {
		return fmt.Errorf("the previous monster file was not found: %s", previousMonster)
	}

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

	var patchList = RenderHeader(sources)
	patchList = append(patchList, removeAllowFromBlock(newList, prevList)...)

	var patchName = strings.Replace(newMonster, ".list", ".update", -1)
	err = writeFile(patchName, []byte(strings.Join(patchList, "\n")))
	if err != nil {
		return err
	}

	fmt.Printf("PATCH: created patch file: %s\n", patchName)
	return nil
}

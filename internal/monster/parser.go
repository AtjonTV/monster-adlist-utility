/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import (
	"fmt"
	"strings"

	"atjon.tv/monster/internal/utils"
)

func (m *Monster) PrepareSourceLists() error {
	for _, source := range m.Sources.Allow {
		err := prepareList(&source)
		if err != nil {
			return err
		}
	}
	for _, source := range m.Sources.Block {
		err := prepareList(&source)
		if err != nil {
			return err
		}
	}
	return nil
}

func prepareList(list *SourceList) error {
	lines, err := utils.ReadLinesFromFile(list.TempFile)
	if err != nil {
		return err
	}

	if list.Trim.Head > 0 {
		lines = lines[list.Trim.Head:]
	}
	if list.Trim.Tail > 0 {
		lines = lines[:len(lines)-list.Trim.Tail]
	}

	trimLines(lines)
	removeComments(lines)

	switch list.Type {
	case "domains":
		break
	case "hosts":
		err := convertHostsToDomains(lines)
		if err != nil {
			return err
		}
	case "abp":
		err := convertABPToBaseDomains(lines)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown list type: '%s' for list '%s'", list.Type, list.Name)
	}

	err = utils.WriteDataToFile(list.TempFile, []byte(strings.Join(lines, "\n")))
	if err != nil {
		return err
	}

	return nil
}

func trimLines(lines []string) {
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
}

func removeComments(lines []string) {
	for i, line := range lines {
		if strings.HasPrefix(line, "#") {
			lines[i] = ""
		}
	}
}

func convertHostsToDomains(lines []string) error {
	for i := range lines {
		if cut, hadPrefix := strings.CutPrefix(lines[i], "0.0.0.0"); hadPrefix {
			lines[i] = strings.TrimSpace(cut)
		}
		if cut, hadPrefix := strings.CutPrefix(lines[i], "127.0.0.1"); hadPrefix {
			lines[i] = strings.TrimSpace(cut)
		}
	}
	return nil
}

func convertABPToBaseDomains(lines []string) error {
	for i := range lines {
		if cut, hadPrefix := strings.CutPrefix(lines[i], "||"); hadPrefix {
			lines[i] = strings.TrimSpace(cut)
		}
		if cut, hadSuffix := strings.CutSuffix(lines[i], "^"); hadSuffix {
			lines[i] = cut
		}
	}
	return nil
}

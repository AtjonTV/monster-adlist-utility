/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import (
	"fmt"
	"os"
	"strings"
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
	lines, err := readListData(list.TempFile)
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
		return fmt.Errorf("Unknown list type: '%s' for list '%s'!", list.Type, list.Name)
	}

	err = writeListData(list, lines)
	if err != nil {
		return err
	}

	return nil
}

func readListData(file string) ([]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(data), "\n"), nil
}

func writeListData(list *SourceList, lines []string) error {
	out, err := os.Create(list.TempFile)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	var data = []byte(strings.Join(lines, "\n"))

	_, err = out.Write(data)
	return err
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

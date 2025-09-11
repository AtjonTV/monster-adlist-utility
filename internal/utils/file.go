/*
 * Copyright (c) 2025 Thomas Obernosterer. Licensed under the EUPL-1.2 or later.
 */

package utils

import (
	"os"
	"strings"
)

func WriteDataToFile(path string, data []byte) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)

	_, err = out.Write(data)
	return err
}

func ReadLinesFromFile(file string) ([]string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(data), "\n"), nil
}

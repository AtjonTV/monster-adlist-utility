/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package monster

import "os"
import "gopkg.in/yaml.v3"

// source.yaml loader

func LoadSourcesFromFile(path string) (Sources, error) {
	var data, err = os.ReadFile(path)
	if err != nil {
		return Sources{}, err
	}

	var sources Sources
	err = yaml.Unmarshal(data, &sources)
	if err != nil {
		return Sources{}, err
	}

	return sources, nil
}

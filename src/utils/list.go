/*
 * Copyright (c) 2025 Thomas Obernosterer, licensed under the EUPL
 */

package utils

func RemoveDuplicatesFromList(elements []string) []string {
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

func RemoveListItems(list []string, itemsToRemove []string) []string {
	elementsToRemove := make(map[string]bool)
	for _, element := range itemsToRemove {
		elementsToRemove[element] = true
	}

	result := make([]string, 0, len(list))

	for _, element := range list {
		if !elementsToRemove[element] && element != "" {
			result = append(result, element)
		}
	}
	return result
}

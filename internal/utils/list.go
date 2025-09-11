/*
 * Copyright (c) 2025 Thomas Obernosterer. Licensed under the EUPL-1.2 or later.
 *
 * SPDX-License-Identifier: EUPL-1.2-or-later
 */

package utils

import "slices"

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

// RemoveListItemsMut removes the itemsToRemove from the list in place.
func RemoveListItemsMut(list *[]string, itemsToRemove *[]string) {
	*list = slices.DeleteFunc(*list, func(element string) bool {
		_, found := slices.BinarySearch(*itemsToRemove, element)
		return found
	})
}

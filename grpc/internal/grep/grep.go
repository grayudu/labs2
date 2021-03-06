// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package grep

import (
	"strings"
)

// WordCount returns the number of occurrence of a word in a line.
func WordCount(word, line string) (count int64) {
	var (
		index int
		l     = strings.ToLower(line)
	)
	for _, b := range []byte(l) {
		if b != word[index] {
			index = 0
			continue
		}
		index++
		if index == len(word) {
			count++
			index = 0
		}
	}

	return
}

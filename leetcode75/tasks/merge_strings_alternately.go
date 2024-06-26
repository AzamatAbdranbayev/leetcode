package tasks

import "strings"

// https://leetcode.com/problems/merge-strings-alternately/?envType=study-plan-v2&envId=leetcode-75
func MergeAlternately(word1 string, word2 string) string {
	maxLen := 0

	if len(word2) >= len(word1) {
		maxLen = len(word2)
	} else {
		maxLen = len(word1)
	}

	var str strings.Builder
	for i := 0; i < maxLen; i++ {
		if i < len(word1) {
			str.WriteString(string(word1[i]))
		}

		if i < len(word2) {
			str.WriteString(string(word2[i]))
		}
	}

	return str.String()
}

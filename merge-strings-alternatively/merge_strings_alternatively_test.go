package main

import "testing"

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		name     string
		word1    string
		word2    string
		expected string
	}{
		{"Both words of equal length", "abc", "pqr", "apbqcr"},
		{"Second word longer than first", "ab", "pqrs", "apbqrs"},
		{"First word longer than second", "abcd", "pq", "apbqcd"},
		{"First word empty", "", "pqrs", "pqrs"},
		{"Second word empty", "abcd", "", "abcd"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MergeAlneternately(tc.word1, tc.word2)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}

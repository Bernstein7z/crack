package chapterOne

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	type testCandidates struct {
		input  string
		result bool
	}

	candidates := []testCandidates{
		{"alan", false},
		{"", true},
		{"aln", true},
		{"alna", false},
		{"1234567890", true},
		{"12345678902", false},
	}

	var output bool
	for _, candidate := range candidates {
		output = IsUnique(candidate.input)
		if output != candidate.result {
			t.Errorf("Output %v not equal to expected %v", candidate.input, candidate.result)
		}
	}
}

func TestCheckPermutation(t *testing.T) {
	type testCandidates struct {
		str1, str2 string
		result     bool
	}

	candidates := []testCandidates{
		{"aba", "baa", true},
		{"aba", "abaa", false},
		{"abc", "abd", false},
		{"", "", false},
		{"1234", "4321", true},
		{"aab", "aba", true},
	}

	var output bool
	for _, candidate := range candidates {
		output = CheckPermutation(candidate.str1, candidate.str2)
		if output != candidate.result {
			t.Errorf("Output from %v and %v not equal to expected %v", candidate.str1, candidate.str2, candidate.result)
		}
	}
}

func TestURLify(t *testing.T) {
	type testCandidate struct {
		raw, formatted string
	}

	candidates := []testCandidate{
		{"Mr John Smith      ", "Mr%20John%20Smith"},
		{"", ""},
		{" Aln", "Aln"},
		{" 23 do 23", "23%20do%2023"},
	}

	var output string
	for _, candidate := range candidates {
		output = URLify(candidate.raw)
		if output != candidate.formatted {
			t.Errorf("output for %v is %v.\nExpected: %v", candidate.raw, output, candidate.formatted)
		}
	}
}

func TestCheckPalindromePermutation(t *testing.T) {
	type testCandidate struct {
		input  string
		result bool
	}

	candidates := []testCandidate{
		{"abbbba", true},
		{"abbbab", true},
		{"abbbbbaa", false},
		{"abcbcb", false},
		{"abcab", true},
		{"abcabab", false},
		{"abcabaabb", true},
		{"0001000", true},
		{"Tact Coa", true},
	}

	var output bool
	for _, candidate := range candidates {
		output = CheckPalindromePermutation(candidate.input)
		if output != candidate.result {
			t.Errorf("output for %v is %v.\nExpected: %v", candidate.input, output, candidate.result)
		}
	}
}

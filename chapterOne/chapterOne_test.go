package chapterOne

import "testing"

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

	for _, candidate := range candidates {
		output := IsUnique(candidate.input)
		if output != candidate.result {
			t.Errorf("Output %v not equal to expected %v", candidate.input, candidate.result)
		}
	}
}

func TestCheckPermutation(t *testing.T) {
	type testCandidates struct {
		str1   string
		str2   string
		result bool
	}

	candidates := []testCandidates{
		{"aba", "baa", true},
		{"aba", "abaa", false},
		{"abc", "abd", false},
		{"", "", false},
		{"1234", "4321", true},
		{"aab", "aba", true},
	}

	for _, candidate := range candidates {
		output := CheckPermutation(candidate.str1, candidate.str2)
		if output != candidate.result {
			t.Errorf("Output from %v and %v not equal to expected %v", candidate.str1, candidate.str2, candidate.result)
		}
	}
}

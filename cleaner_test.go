package main

import "testing"

func TestCleanOfProfanity(t *testing.T) {
	tests := []struct {
		name string
		input string
		expected string
	} {
		{
			name: "return untouched",
			input: "I had something interesting for breakfast",
			expected: "I had something interesting for breakfast",
		},
		{
			name: "remove sharbert",
			input: "I hear Mastodon is better than Chirpy. sharbert I need to migrate",
			expected: "I hear Mastodon is better than Chirpy. **** I need to migrate",
		},
		{
			name: "Remove kerfuffle and Fornax",
			input: "I really need a kerfuffle to go to bed sooner, Fornax !",
			expected: "I really need a **** to go to bed sooner, **** !",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := cleanOfProfanity(tc.input)

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected string: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
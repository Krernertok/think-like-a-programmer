package words

import "testing"

type filterTest struct {
	letters           []rune
	include           bool
	expectedWordCount int
}

func TestFilterWords(t *testing.T) {
	words := []string{
		"hello",
		"world",
	}

	lettersInBothWords := []rune{
		'l',
		'o',
	}

	lettersInNeitherWord := []rune{
		'a',
		'x',
	}

	lettersInOneWord := []rune{
		'e',
		'r',
	}

	testCases := []filterTest{
		filterTest{
			letters:           lettersInBothWords,
			include:           true,
			expectedWordCount: 2,
		},
		filterTest{
			letters:           lettersInNeitherWord,
			include:           true,
			expectedWordCount: 0,
		},
		filterTest{
			letters:           lettersInOneWord,
			include:           true,
			expectedWordCount: 1,
		},
		filterTest{
			letters:           lettersInBothWords,
			include:           false,
			expectedWordCount: 0,
		},
		filterTest{
			letters:           lettersInNeitherWord,
			include:           false,
			expectedWordCount: 2,
		},
		filterTest{
			letters:           lettersInOneWord,
			include:           false,
			expectedWordCount: 1,
		},
	}

	for _, testCase := range testCases {
		for _, letter := range testCase.letters {
			filteredWords := FilterWords(words, letter, testCase.include)
			numWords := len(filteredWords)

			if len(filteredWords) != testCase.expectedWordCount {
				t.Error(
					"Expected length of", testCase.expectedWordCount,
					"got:", numWords,
				)
			}
		}
	}
}

func TestGetPattern(t *testing.T) {
	s := "soliloquy"
	letter := 'l'
	expected := []rune{0, 0, 'l', 0, 'l', 0, 0, 0, 0}
	pattern := GetPattern(s, letter)

	if !patternsMatch(pattern, expected) {
		t.Error(
			"Expected pattern:", expected,
			"Got pattern:", pattern,
		)
	}
}

func TestMatchPattern(t *testing.T) {
	s := "soliloquy"
	letter := 'l'
	pattern1 := []rune{0, 0, 'l', 0, 'l', 0, 0, 0, 0}
	pattern2 := []rune{0, 0, 'l', 0, 'l', 0, 0, 0, 'l'}

	if !matchPattern(s, pattern1, letter) {
		t.Error("Expected", []rune(s), "to match pattern", pattern1)
	}

	if matchPattern(s, pattern2, letter) {
		t.Error("Did not expect", []rune(s), "to match pattern", pattern2)
	}
}

func TestFilterOutPattern(t *testing.T) {
	originalWords := []string{
		"wenlock",
		"wetland",
		"whalery",
		"westlan",
		"whiglet",
		"whamble",
		"whetile",
		"whiffle",
		"wendell",
		"wenchel",
	}

	expectedWords := []string{
		"whamble",
		"whetile",
		"whiffle",
	}

	letter := 'l'
	pattern := []rune{0, 0, 0, 0, 0, 'l', 0}

	words := filterPattern(originalWords, pattern, letter, true)

	if len(words) != len(expectedWords) {
		t.Error(
			"Expected:", expectedWords,
			"Got:", words,
		)
	} else {
		for _, word := range expectedWords {
			if !stringInSlice(word, words) {
				t.Error(
					"Expected:", expectedWords,
					"Got:", words,
				)
				break
			}
		}
	}
}

func TestGetCommonestPattern(t *testing.T) {
	letter := 'l'
	words := []string{
		"wenlock",
		"wetland",
		"whalery",
		"westlan",
		"whiglet",
		"whamble",
		"whetile",
		"whiffle",
		"wendell",
		"wenchel",
	}

	expectedWords := []string{
		"wenlock",
		"wetland",
		"whalery",
	}

	expectedPattern := []rune{0, 0, 0, 'l', 0, 0, 0}

	wordlist, pattern := getCommonestPattern(words, letter)

	if len(wordlist) != len(expectedWords) {
		t.Error(
			"Expected:", expectedWords,
			"Got:", wordlist,
		)
	} else {
		for _, word := range expectedWords {
			if !stringInSlice(word, wordlist) {
				t.Error(
					"Expected:", expectedWords,
					"Got:", wordlist,
				)
				break
			}
		}
	}

	if !patternsMatch(pattern, expectedPattern) {
		t.Error(
			"Expected:", expectedPattern,
			"Got:", pattern,
		)
	}
}

func stringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if str == s {
			return true
		}
	}

	return false
}

func patternsMatch(p1, p2 []rune) bool {
	if len(p1) != len(p2) {
		return false
	}

	for i, r := range p1 {
		if r != p2[i] {
			return false
		}
	}

	return true
}

func TestGetUpdatedList(t *testing.T) {
	words := []string{
		"wenlock",
		"wetland",
		"whalery",
		"westlan",
		"whiglet",
		"whamble",
		"whetile",
		"whiffle",
		"wendell",
		"wenchel",
	}

	expectedWords := []string{
		"whalery",
		"whiglet",
		"whamble",
		"whetile",
		"whiffle",
	}

	letter1 := 'x'
	letter2 := 'h'

	updatedWords1, pattern1 := GetUpdatedList(words, letter1)
	updatedWords2, pattern2 := GetUpdatedList(words, letter2)

	expectedPattern1 := []rune{}
	expectedPattern2 := []rune{0, 'h', 0, 0, 0, 0, 0}

	if !patternsMatch(pattern1, expectedPattern1) {
		callError(t, expectedPattern1, pattern1)
	}

	if len(updatedWords1) != len(words) {
		callError(t, words, updatedWords1)
	}

	for _, word := range words {
		if !stringInSlice(word, updatedWords1) {
			callError(t, words, updatedWords1)
		}
	}

	if !patternsMatch(pattern2, expectedPattern2) {
		callError(t, expectedPattern2, pattern2)
	}

	if len(updatedWords2) != len(expectedWords) {
		callError(t, expectedWords, updatedWords2)
	}

	for _, word := range expectedWords {
		if !stringInSlice(word, updatedWords2) {
			callError(t, expectedWords, updatedWords2)
		}
	}

}

func TestSamePattern(t *testing.T) {
	p1 := []rune{}
	p2 := []rune{0, 0, 0, 'a', 0, 0, 0}

	if !SamePattern(p1, p1) {
		t.Error("A pattern should be the same as itself.")
	}

	if !SamePattern(p2, p2) {
		t.Error("A pattern should be the same as itself.")
	}

	if SamePattern(p1, p2) {
		t.Error("Different patterns should not be considered the same.")
	}
}

func TestMergePatterns(t *testing.T) {
	p1 := []rune{'a', 'b', 'c', 0, 0}
	p2 := []rune{0, 0, 0, 'x', 'y'}
	expected := []rune{'a', 'b', 'c', 'x', 'y'}
	mergedPattern := MergePatterns(p1, p2)

	if !SamePattern(mergedPattern, expected) {
		callError(t, expected, mergedPattern)
	}
}

func callError(t *testing.T, expected, got interface{}) {
	t.Error(
		"Expected:", expected,
		"Got:", got,
	)
}

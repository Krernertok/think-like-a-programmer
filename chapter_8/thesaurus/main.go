package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Thesaurus represents a map with string keys and slice of string values
type Thesaurus map[string][]string

type synonym struct {
	word      string
	linkCount int
}

func (t Thesaurus) getSynonyms(word string) []string {
	ranking := make(map[string]int)

	synonyms := t[word]

	for _, synonym := range synonyms {
		ranking[synonym]++
		secondarySynonyms := t[synonym]
		for _, ss := range secondarySynonyms {
			ranking[ss]++
		}
	}

	rankedSynonyms := []synonym{}
	for word, count := range ranking {
		rankedSynonyms = append(rankedSynonyms, synonym{word, count})
	}
	sort.Sort(byLinkCount(rankedSynonyms))

	result := []string{}
	for _, s := range rankedSynonyms {
		result = append(result, s.word)
	}

	return result
}

type byLinkCount []synonym

func (b byLinkCount) Len() int {
	return len(b)
}

func (b byLinkCount) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byLinkCount) Less(i, j int) bool {
	return b[i].linkCount > b[j].linkCount
}

func generateThesaurus() Thesaurus {
	thesaurus := make(map[string][]string)

	wordfile, err := os.Open("data/thesaurus.txt")
	defer wordfile.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(wordfile)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, ",")
		thesaurus[words[0]] = words[1:]
	}

	return Thesaurus(thesaurus)
}

func getWord() string {
	var word string
	args := os.Args[1:]
	numArgs := len(args)

	if numArgs == 0 {
		fmt.Println("Please provide a word as a argument, e.g. 'thesaurus book'")
		os.Exit(0)
	}

	if numArgs == 1 {
		word = args[0]
	} else {
		word = strings.Join(args, " ")
	}

	return word
}

func main() {
	thesaurus := generateThesaurus()
	word := getWord()
	synonyms := thesaurus.getSynonyms(word)

	if len(synonyms) > 5 {
		fmt.Println(synonyms[1:6])
	} else if len(synonyms) == 2 {
		fmt.Println(synonyms[1])
	} else {
		fmt.Println("No synonyms found.")
	}
}

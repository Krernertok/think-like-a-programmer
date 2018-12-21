package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	thesaurus := generateThesaurus()
	word := "book"
	synonyms := thesaurus.getSynonyms(word)
	fmt.Println(synonyms[:10])
	fmt.Println(thesaurus[word])

	// read command argument(s) which should be a word or phrase
	// return synonyms for the word/phrase
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

// Thesaurus represents a map with string keys and slice of string values
type Thesaurus map[string][]string

type synonym struct {
	word      string
	linkCount int
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

func (t Thesaurus) getSynonyms(word string) []string {
	ranking := make(map[string]int)
	t.rankSynonyms(word, &ranking)

	synonyms := []synonym{}
	for k, v := range ranking {
		synonyms = append(synonyms, synonym{k, v})
	}
	sort.Sort(byLinkCount(synonyms))

	result := []string{}

	for _, s := range synonyms {
		result = append(result, s.word)
	}

	return result
}

func (t Thesaurus) rankSynonyms(word string, ranking *map[string]int) {
	_, ranked := (*ranking)[word]

	if !ranked {
		(*ranking)[word] = 1
	}

	synonyms, found := t[word]
	if !found {
		return
	}

	for _, synonym := range synonyms {
		_, keyExists := (*ranking)[synonym]
		if !keyExists {
			t.rankSynonyms(synonym, ranking)
		}
		(*ranking)[synonym]++
	}
}

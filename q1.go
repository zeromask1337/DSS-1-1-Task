package cos418_hw1_1

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

// Find the top K most common words in a text document.
// 	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuations and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func topWords(pathString string, numWords int, charThreshold int) []WordCount {
	contents, _ := ioutil.ReadFile(pathString)
	loweredText := strings.ToLower(string(contents))

	var punctuation = regexp.MustCompile(`[[:punct:]]`)
	filteredText := punctuation.ReplaceAllString(loweredText, "")
	wordList := strings.Fields(filteredText)

	counts := make([]WordCount, 0)
	//counts = append(counts, WordCount{"gay", 1})
	for _, word := range wordList {
		//counts = append(counts, WordCount{word, 1})
		if len(word) >= charThreshold {
			counts = append(counts, WordCount{word, 1})

			gotIt := 0
			notGotIt :=0
			for _, draw := range wordList {
				//counts[i].Word = word

				//fmt.Println("Draw is: ", draw)

				if draw == word {
						//counts[i].Count += 1
					gotIt += 1
				} else {
						//counts[i].Count = 1
					notGotIt += 1
				}
				//TODO: отслеживание работает, осталось насадить сову на глобус. Перед продолжением желательно запустить скрипт.
				//ok := counts[i].Word
				//if ok == word {
				//	counts[i].Count += 1
				//} else {
				//	counts[i].Count = 1
				//}
			}
					fmt.Println(gotIt, notGotIt)
		}
	}
	fmt.Println("Wordlist: ", wordList,"\nCounts: " , counts)
	return nil
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}
// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

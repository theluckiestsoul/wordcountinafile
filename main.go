package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var wordStore map[string]int
var sortedWords []wc

func main() {
	wordStore = map[string]int{}
	sortedWords = make([]wc, 0)
	fs, _ := os.OpenFile("abc.txt", os.O_CREATE, os.ModeAppend)

	defer fs.Close()
	reader := bufio.NewReader(fs)
	for {
		if _, err := reader.Peek(1); err != nil {
			break
		}
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
		}
		storeWords(string(line))
	}
	arrangeWordsInDsc()
	fmt.Println(sortedWords)
}
func storeWords(line string) {
	words := strings.Split(line, " ")
	for _, word := range words {
		wordStore[word] = wordStore[word] + 1
	}
}

type wc struct {
	word  string
	count int
}

func arrangeWordsInDsc() {
	slice := make([]wc, 0)
	for key, val := range wordStore {
		slice = append(slice, wc{word: key, count: val})
	}
	sort.Sort(byLength(slice))
	sortedWords = append(sortedWords, slice...)
}

type byLength []wc

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return s[i].count > s[j].count
}

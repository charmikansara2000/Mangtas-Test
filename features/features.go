package features

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/ledongthuc/pdf"
)

func ReadPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	AnsString := reg.ReplaceAllString(buf.String(), " ")
	return AnsString, nil
}

//Finds the number of occurance of each word and returns the map[string]int
func WordCount(content string) map[string]int {

	pdfContent := strings.Fields(content)

	totalWords := make(map[string]int)

	for _, word := range pdfContent {
		_, wordExists := totalWords[word]
		if wordExists {
			totalWords[word] += 1
		} else {
			totalWords[word] = 1
		}
	}

	return totalWords
}

//Sorts the recieved map by the value and prints top 10 occuring words
func PrintAnswer(input map[string]int) {
	newMap := map[int][]string{}
	var a []int
	for key, val := range input {
		newMap[val] = append(newMap[val], key)

	}
	for key := range newMap {
		a = append(a, key)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	limit := 0
	for _, key := range a {
		for _, s := range newMap[key] {
			if limit < 10 {
				fmt.Printf("%s -> %d times\n", s, key)
				limit++
			}
		}
	}
}

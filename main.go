package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/go-rod/rod"
	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <URL>")
		return
	}

	url := os.Args[1]

	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Invalid URL:", err)
		return
	}

	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(url).MustWaitIdle()

	html := page.MustHTML()

	text := extractTextFromHTML(html)

	re := regexp.MustCompile(`\b\w+\b`)
	words := re.FindAllString(text, -1)

	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[strings.ToLower(word)]++
	}

	type wordCountPair struct {
		word  string
		count int
	}
	var wordCountPairs []wordCountPair
	for word, count := range wordCounts {
		wordCountPairs = append(wordCountPairs, wordCountPair{word, count})
	}
	sort.Slice(wordCountPairs, func(i, j int) bool {
		return wordCountPairs[i].count > wordCountPairs[j].count
	})

	for _, pair := range wordCountPairs {
		fmt.Println(pair.word)
	}
}

func extractTextFromHTML(htmlStr string) string {
	var text string
	tokenizer := html.NewTokenizer(strings.NewReader(htmlStr))

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			fmt.Println("Error tokenizing HTML:", err)
			return ""
		}

		if tokenType == html.TextToken {
			text += string(tokenizer.Text())
		}
	}

	return text
}

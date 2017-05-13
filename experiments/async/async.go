package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	urls := []string{
		"https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=en",
		"https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=zh-cn",
		"https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=de",
		"https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity/transcript?language=ru",
	}

	for _, url := range urls {

		//		fmt.Println(talkTexts(url))

		//fmt.Println( <-funcWithChanResult(url))
		fmt.Println(funcWithNonChanResult(url))
	}

	//var texts []string

	//fmt.Println(texts)
}

func funcWithChanResult(url string) chan []string {
	chanStr := make(chan []string)
	go func() {
		chanStr <- talkTexts(url)
	}()
	return chanStr
}

func funcWithNonChanResult(url string) []string {
	return <-funcWithChanResult(url)
}

func talkTexts(url string) []string {
	//textChannel := make(chan []string)
	var para []string
	page, _ := goquery.NewDocument(url)

	texts := page.Find(".talk-transcript__para__text").Contents().Text()

	for _, text := range strings.Split(texts, "\n\n") {

		//fmt.Println(text)
		para = append(para, text)

	}
	return para
}

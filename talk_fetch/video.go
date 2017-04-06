package talk_fetch

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func speaker(doc *goquery.Document) {
	speaker := doc.Find(".talk-speaker__name").Contents().Text()
	fmt.Println(speaker)
}

func title(doc *goquery.Document) {
	title := doc.Find(".player-hero__title__content").Contents().Text()
	fmt.Println(title)
}

func duration(doc *goquery.Document) {

	duration := doc.Find(".player-hero__meta").Contents().Text()
	//fmt.Println(duration)

	//for _, x := range strings.Split(duration, "\n") {
	//	fmt.Println(x)
	//	println("~~~~~~")
	//}

	x := strings.Split(duration, "\n")
	fmt.Println(x[6])

}

func time_filmed(doc *goquery.Document) {

	time_filmed := doc.Find(".player-hero__meta").Contents().Text()

	//	fmt.Println(time_filmed)

	y := strings.Split(time_filmed, "\n")
	fmt.Println(y[11])

}

func talk_views_count(doc *goquery.Document) {

	talk_views_count := doc.Find("#sharing-count").Contents().Text()
	//	fmt.Println(talk_views_count)

	a := strings.Split(talk_views_count, "\n")
	b := strings.TrimSpace(a[2])
	fmt.Println(b)

}

func talk_topics_list(doc *goquery.Document) {

	talk_topics := doc.Find(".talk-topics__list").Contents().Text()

	c := strings.Split(talk_topics, "\n")

	for i := 3; i < len(c); i++ {
		fmt.Println(c[i])
	}

}

func talk_comments_count(doc *goquery.Document) {

	talk_comments_count := doc.Find(".h11").Contents().Text()
	//fmt.Println(talk_comments_count)
	d := strings.Split(talk_comments_count, " ")
	fmt.Println(d[0])

}

func main() {

	//	url := "https://www.ted.com/talks/ken_robinson_says_schools_kill_creativity"
	url := "http://www.ted.com/talks/simon_sinek_how_great_leaders_inspire_action"

	/*
		var url string

		fmt.Scanf("%s", &url)

	*/

	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}

	//speaker(doc)
	//title(doc)
	//duration(doc)
	//time_filmed(doc)
	//talk_views_count(doc)
	//talk_topics_list(doc)
	//talk_comments_count(doc)

}
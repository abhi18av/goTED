package main

import (
	"encoding/json"
	"fmt"
//	"io/ioutil"
//	"os"
	"strings"

//	"github.com/Jeffail/gabs"
	"github.com/PuerkitoBio/goquery"
)

var langCodes = map[string]string{
	"Chinese, Simplified": "zh-cn",
	"English":             "en",
	"German":              "de",
	"Russian":             "ru",
}

type kenJSON struct {
	// This is the main title from English language
	TalkTitle string `json:"talk_title"`
	// The link to the main video
	TalkLink string `json:"talk_link"`
/*
	// The main container for the Transcript text
	Transcript struct {
		En   []string `json:"en"`
		ZhCn []string `json:"zh-cn"`
		De   []string `json:"de"`
		Ru   []string `json:"ru"`
	} `json:"transcript"`
*/
}

// From the main page or the English language transcript
func title(doc *goquery.Document) {
	title := doc.Find(".player-hero__title__content").Contents().Text()
	fmt.Println(title)
}

func texts(doc *goquery.Document) {
	texts := doc.Find(".talk-transcript__para__text").Contents().Text()
	for _, text := range strings.Split(texts, "  ") {

		fmt.Println(text)
	}
}

func (p kenJSON) toString() string {
	return toJSON(p)
}

func toJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	return string(bytes)
}

// this should return an array of strings => ["langs"]
func langs(doc *goquery.Document) []string {

	var langsList []string

	langs := doc.Find(".talk-transcript__language").Contents().Text()

	//fmt.Println(langs)
	langsSeparated := strings.Split(langs, "\n")

	for i := 1; i < len(langsSeparated)-1; i++ {
		//fmt.Println(i, ":", langsSeparated[i])
		langsList = append(langsList, langsSeparated[i])
	}

	return langsList
}

func main() {

	var doc []*goquery.Document

	url := "https://www.ted.com/talks/ari_wallach_3_ways_to_plan_for_the_very_long_term/transcript?language=en"
//	mainURL := "https://www.ted.com/talks/ari_wallach_3_ways_to_plan_for_the_very_long_term"


	page, _ := goquery.NewDocument(url)
	doc = append(doc, page)

if texts(page) == ""{
	print("Not available yet")
}


	//fmt.Println(doc[0])

	//fmt.Println(langs(doc)[1])

/*
	
			var availableLangs []string

			for _, code := range langs(doc[0]) {

				if langCodes[code] == "" {
					//fmt.Println(code, " : Not available")
				} else {
					//fmt.Println(langCodes[code])
					availableLangs = append(availableLangs, langCodes[code])
				}
			}

		fmt.Println(availableLangs)
	
for _,x := range availableLangs{
	transcriptURL := mainURL + "/transcript?language=" + x

	// asynchronously send the requests and fill in the correct place in the struct
	fmt.Println(transcriptURL)
}

*/



/*
	raw, _ := ioutil.ReadFile("../resources/al-A_simple_way_to_break_a_bad_habit.json")
	
		var TED_JSONs []TED_JSON
		json.Unmarshal(raw, &TED_JSONs)

		for _, p := range TED_JSONs {
			fmt.Println(p.toString())
		}

		fmt.Println(toJson(TED_JSONs))
	
	jsonParsed, _ := gabs.ParseJSON(raw)

	x, _ := jsonParsed.Path("transcript.zh-cn").Children()

	for _, child := range x {

		fmt.Println(child.Data())
	}

*/

}
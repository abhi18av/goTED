package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type alphanumeric struct {
	Alphabet string            `json:"anAlphabet"`
	Number   string            `json:"aNumber"`
	AlNum    map[string]string `json:"AlNum"`
}

func (someStruct alphanumeric) pairAlphanumeric() string {

	return someStruct.Number + someStruct.Alphabet

}

func main() {

	var wg sync.WaitGroup

	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	alphabets := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	//var aleph alphanumeric
	var alephS []alphanumeric
	n := 10 // number of codes you want to print

	wg.Add(n)

	z := make(map[string]string)

	for i := 0; i < n; i++ {
		go func(numbers []string, alphabets []string, z map[string]string) {
			defer wg.Done()
			x, y, alp := makeAleph(numbers, alphabets)
			z[y] = x
			alp.AlNum = z
			alephS = append(alephS, alp)
			//fmt.Println(x)
		}(numbers, alphabets, z)
	}

	wg.Wait()
	//fmt.Println(alephS)
	//fmt.Println(z)
	//alephS.AlNum = z
	body, _ := json.Marshal(alephS)
	fmt.Println(string(body))

} // end of main()

func makeAleph(numbers []string, alphabets []string) (string, string, alphanumeric) {

	var aleph alphanumeric

	anAlphabet := aNum(numbers)
	aNumber := anAlph(alphabets)

	aleph.Alphabet = anAlphabet
	aleph.Number = aNumber
	//fmt.Println(aleph.pairAlphanumeric())

	//return aleph.pairAlphanumeric()
	return aNumber, anAlphabet, aleph
}

func randomIndex() int {
	randTime := time.Time.UnixNano(time.Now())

	rand.Seed(randTime)

	return rand.Intn(10)
}

func aNum(numbers []string) string {

	return numbers[randomIndex()]

}

func anAlph(alphabets []string) string {

	return alphabets[randomIndex()]

}

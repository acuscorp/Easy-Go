package once

import (
	"fmt"
	"sync"
)

type SlowComplicatedParser interface {
	Parse(string) string
}

var parser SlowComplicatedParser

var once sync.Once

func InitOnceParser() {
	fmt.Printf("Parse(\"Love of my life\"): %v\n", Parse("Love of my life"))
	fmt.Printf("Parse(\"The man who sold the world\"): %v\n", Parse("The man who sold the world"))
}
func Parse(dataToParse string) string {
	once.Do(func() {
		parser = InitParser()
	})
	return parser.Parse(dataToParse)
}

var parserCounter = 0

func InitParser() SlowComplicatedParser {
	parserCounter++
	fmt.Println("Init parsed called ", parserCounter)
	return JSON{}
}

type JSON struct{}

func (json JSON) Parse(data string) string {
	return fmt.Sprintf("this is simple song: %s ", data)
}

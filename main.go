package main

import (
  "fmt"
	"main/poetry"
)

func main() {
	p := poetry.Poem{{"It will be soon midsummer of a fleeting year,",
		"The sward will be brown, flowers faint and few,",
		"Songbirds are hushed all but a faint but clear song",
		"And 'larum of the bird-boy reach the ear,",
		"Through the warm air floats forth lime's sweet scent,",
		"And wayside branches have lost the rose's bloom.",
		"The corn is golden along a thousand sea like slopes"},
	}
	v, c := p.Stats()
  fmt.Printf("Vowels: %d, Consonants: %d\n",v, c)
  fmt.Printf("Stanzas: %d Lines: %d\n", p.NumStanzas(),p.NumLines() )
}

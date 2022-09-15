package poetry

import (
  "testing"
)

func TestNumStanzas(t *testing.T) {
  p := Poem{}
  p = Poem{{"It will be soon midsummer of a fleeting year,",
		"The sward will be brown, flowers faint and few,",
		"Songbirds are hushed all but a faint but clear song",
		"And 'larum of the bird-boy reach the ear,",
		"Through the warm air floats forth lime's sweet scent,",
		"And wayside branches have lost the rose's bloom.",
		"The corn is golden along a thousand sea like slopes"},
	}

  if p.NumStanzas() != 7 {
    t.Fatalf("unexpected stanza count %d", p.NumStanzas())
  }
}

func TestStats(t *testing.T){
  p:= Poem{}

  v,c := p.Stats()
  if v != 0 || c != 0 {
    t.Fatalf("Bad number of vowels or consonants")
  }
}
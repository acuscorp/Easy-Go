package serializing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func InitJsonArray() {
	const data = `
		{"name": "Fred", "age": 40}
		{"name": "Mary", "age": 21}
		{"name": "Pat", "age": 30}
	`
  fmt.Println("data: ", data)
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	dec := json.NewDecoder(strings.NewReader(data))
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
  fmt.Println("Json encoding output and saved in a structure")
	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		fmt.Println(t)
		err = enc.Encode(t)
		if err != nil {
			panic(err)
		}
	}
  fmt.Println("Json decoding output bytes to json")
	out := b.String()
	fmt.Println(out)
}

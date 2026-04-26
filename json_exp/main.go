package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println(json.Valid(nil))
	fmt.Println(json.Valid([]byte("")))
	fmt.Println(json.Valid([]byte{}))
	fmt.Println(json.Valid([]byte("{}")))

	var v []byte
	fmt.Println(v == nil)
	fmt.Println(len(v))
	fmt.Println(len([]byte{}))
}

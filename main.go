package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
)

var isChar = regexp.MustCompile(`^[a-zA-Z]+$`)

func main() {
	for _, r := range os.Args[1] {
		c := string(r)
		if !isChar.MatchString(c) {
			s := hex.EncodeToString([]byte(c))
			fmt.Printf("%%%s", s)
		} else {
			fmt.Print(c)
		}
	}
	fmt.Print("\n")
}

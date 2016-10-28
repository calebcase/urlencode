package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	u, err := url.Parse(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(u.String())
}

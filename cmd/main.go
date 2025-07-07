package main

import (
	"fmt"
	"log"

	"github.com/escabora/parser-ua/internal/parser"
)

func main() {
	p, err := parser.NewParser("testdata/browscap.csv")
	if err != nil {
		log.Fatal(err)
	}

	ua := "Mozilla/5.0 (Linux; Android 10; Redmi Note 8 Pro Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/137.0.7151.117 Mobile Safari/537.36 (Mobile; afma-sdk-a-v250505999.243220000.1)"
	r := p.Parse(ua)
	if r != nil {
		fmt.Printf("Browser: %s %s\nPlatform: %s\nDevice: %s\nPattern: %s\n",
			r.Browser, r.Version, r.Platform, r.DeviceType, r.Matched)
	} else {
		fmt.Println("UA not recognized")
	}
}

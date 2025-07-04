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

	r := p.Parse("Mozilla/5.0 (Windows NT 10.0; Win64; x64)...")
	if r != nil {
		fmt.Printf("Browser: %s %s\nPlatform: %s\nDevice: %s\nPattern: %s\n",
			r.Browser, r.Version, r.Platform, r.DeviceType, r.Matched)
	} else {
		fmt.Println("UA not recognized")
	}
}

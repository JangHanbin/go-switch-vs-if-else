package main

import (
	"log"
	"time"
)

func main() {

	cfg := "aws"
	n := 0
	start := time.Now()
	for n < 1000000000 {
		if cfg == "aws" {
			//log.Printf("AWS ")
			//log.Print(string(n) + " ")
		} else if cfg == "Azure" {
			//log.Printf("Azure")
		} else {
			//log.Printf("Good")
		}
		n += 1
	}

	elapsed := time.Since(start)
	log.Printf("excution tooks %s", elapsed)
	n = 1
	start = time.Now()
	for n < 1000000000 {
		switch cfg {
		case "aws":
			//log.Print(string(n) + " ")
			//log.Fatalf("AWS")
		case "Azure":
			//log.Printf("Azure")
		default:
			//log.Printf("Good")
		}
		n += 1
	}
	elapsed = time.Since(start)
	log.Printf("excution tooks %s", elapsed)
}

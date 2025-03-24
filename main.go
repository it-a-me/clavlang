package main

import (
	"bufio"
	"log"
	"os"

	"github.com/it-a-me/clavlang/scanner"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("-- ")
	if len(os.Args) > 2 {
		log.Fatal("Please supply 0-1 file arguments")
	}
	if len(os.Args) == 2 {
		runFile(os.Args[1])
	}
	repl()
}

func repl() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	for err == nil {
		run(text)
		text, err = reader.ReadString('\n')
	}
	log.Fatal(err)
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	run(string(bytes))
}

func run(text string) {
	s := scanner.NewScanner(text)

	tokens, errs := s.Scan()
	for _, t := range tokens {
		log.Println(t.String())
	}
	if errs != nil {
		for _, err := range errs {
			log.Print(err.String())
		}
		os.Exit(1)
	}
}

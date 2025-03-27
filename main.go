package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/it-a-me/clavlang/parser"
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
	if errs != nil {
		for _, err := range errs {
			log.Print(err)
		}
		os.Exit(1)
	}
	log.Print("[")
	for _, t := range tokens {
		log.Printf(" %s", t.Type.String())
	}
	log.Println("]")

	p := parser.NewParser(tokens)
	expr, errs := p.Parse()
	if errs != nil {
		for _, err := range errs {
			log.Print(err)
		}
		os.Exit(1)
	}
	fmt.Println(parser.AstString(expr))
}

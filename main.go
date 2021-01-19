package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lenguti/STLParser/parser"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("main: unable to parse file argument [%v]", os.Args)
	}

	fileArg := os.Args[1]
	f, err := os.OpenFile(fileArg, os.O_RDONLY, 0755)
	defer f.Close()
	if err != nil {
		log.Fatalf("main: unable to open file [%s]", err)
	}

	p := parser.New(f)
	s, err := p.Parse()
	if err != nil {
		log.Fatalf("main: unable to parse file [%s]", err)
	}

	fmt.Println("PARSED SOLID")
	fmt.Println(s)
	fmt.Printf("SOLID HAS %d Triangles\n", len(s.Facets))
	fmt.Printf("SOLID SURFACE AREA: [%f]\n", s.SurfaceArea())
}

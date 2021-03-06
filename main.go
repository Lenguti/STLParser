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

	if s.CheckDuplicates() {
		log.Fatalf("main: stl file has duplicate triangles")
	}

	min, max := s.BoundingBox()
	fmt.Printf("Number of triangles: %d\n", len(s.Facets))
	fmt.Printf("Surface area       : %f\n", s.SurfaceArea())
	fmt.Printf("Bounding box       : %+v %+v\n", min, max)
}

/*
	We want a new feature to return an error when we receive any duplicated triangles.

	- After parsing the file we can check and see if any triangles are duplicated from
	  going over the solids triangles.
*/

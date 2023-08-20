package main

import (
	"graph-docs-golang/lib"
)

func main() {
	parser := lib.Parse("./test", "output", "graph")

	parser.Generate()
}

package main

import (
	"graph-docs-golang/lib"
)

func main() {
	parser := lib.Parse("./test")

	parser.Generate()
}

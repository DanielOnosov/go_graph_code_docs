package main

import (
	"github.com/DaksinWorld/go_graph_code_docs"
	"github.com/DaksinWorld/go_graph_code_docs/themes"
	"log"
)

func main() {
	parser := docParser.Parse("./app", "output", "graph")

	parser.AddTitle("Process of node interaction")

	parser.AddTheme(themes.LightTheme)

	err := parser.Generate()

	log.Fatal(err)
}

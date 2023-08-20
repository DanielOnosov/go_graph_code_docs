package main

import (
	"github.com/DaksinWorld/go_graph_code_docs"
	"github.com/DaksinWorld/go_graph_code_docs/themes"
)

func main() {
	parser := docParser.Parse("./app", "output", "graph")

	parser.AddTitle("Process of node interaction")

	parser.AddTheme(themes.DarkTheme)

	parser.Generate()
}

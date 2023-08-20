package main

func main() {
	parser := lib.Parse("./test", "output", "graph")

	parser.Generate()
}

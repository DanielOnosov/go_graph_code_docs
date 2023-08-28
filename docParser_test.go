package docParser

import (
	"github.com/DaksinWorld/go_graph_code_docs/themes"
	"os"
	"testing"
)

func TestParser_Generate(t *testing.T) {
	path := "./example"
	parser := Parse(path+"/app", path+"/output", "graph")

	parser.AddTitle("Process of node interaction")

	parser.AddTheme(themes.LightTheme)

	err := parser.Generate()

	if err != nil {
		t.Error("Error while generating output: " + err.Error())
	}

	_, err = os.Stat(path + "/app")
	if err != nil {
		t.Error("Error after generating output: " + err.Error())
	}
}

func BenchmarkParser_Generate(b *testing.B) {
	path := "./example"
	parser := Parse(path+"/app", path+"/output", "graph")

	parser.AddTitle("Process of node interaction")

	parser.AddTheme(themes.LightTheme)

	_ = parser.Generate()
}

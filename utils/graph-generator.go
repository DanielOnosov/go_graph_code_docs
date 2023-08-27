package utils

import (
	"bytes"
	"fmt"
	"github.com/DaksinWorld/go_graph_code_docs/structs"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/fatih/color"
	"github.com/goccy/go-graphviz"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GenerateChart(nodes []structs.Node, edges []structs.Edge, outputFolder string, outputName string, title string, theme structs.Theme) error {
	g := graph.New(graph.StringHash, graph.Directed())
	var titleOfGraph string

	if title == "" {
		titleOfGraph = "My Graph"
	} else {
		titleOfGraph = title
	}

	for _, node := range nodes {
		var attrs []func(properties *graph.VertexProperties)

		if len(theme.CollOfAttrs) >= 1 {
			for _, a := range theme.CollOfAttrs[0] {
				attrs = append(attrs, graph.VertexAttribute(a.Key, a.Value))
			}
		}

		for _, attr := range node.Attributes {
			if strings.ToLower(attr.Key) == "description" {
				node.Label += "\n" + attr.Value
			} else {
				attrs = append(attrs, graph.VertexAttribute(attr.Key, attr.Value))
			}
		}

		attrs = append(attrs, graph.VertexAttribute("label", node.Label))

		_ = g.AddVertex(node.Id, attrs...)
	}

	for _, edge := range edges {
		var attrs []func(*graph.EdgeProperties)

		if len(theme.CollOfAttrs) >= 2 {
			for _, a := range theme.CollOfAttrs[1] {
				attrs = append(attrs, graph.EdgeAttribute(a.Key, a.Value))
			}
		}

		for _, attr := range edge.Attributes {
			attrs = append(attrs, graph.EdgeAttribute(attr.Key, attr.Value))
		}

		_ = g.AddEdge(edge.From, edge.To, attrs...)
	}

	CreateOutputFolder(outputFolder)

	file, err := os.Create(fmt.Sprintf("./%s/", outputFolder) + outputName + ".gv")

	if err != nil {
		fmt.Println(err.Error())
	}

	var bgcolor = "#ffffff"
	var fontcolor = "#111111"

	if theme.Name == "DarkTheme" {
		bgcolor = "#111111"
		fontcolor = "#ffffff"
	}

	err = draw.DOT(g, file,
		draw.GraphAttribute("rankdir", "LR"),
		draw.GraphAttribute("labelloc", "top"),
		draw.GraphAttribute("fontname", "Arial"),
		draw.GraphAttribute("bgcolor", bgcolor),
		draw.GraphAttribute("fontcolor", fontcolor),
		draw.GraphAttribute("label", titleOfGraph),
	)

	b, _ := ioutil.ReadFile(fmt.Sprintf("./%s/", outputFolder) + outputName + ".gv")

	graphV, err := graphviz.ParseBytes(b)
	gT := graphviz.New()

	image, err := gT.RenderImage(graphV)
	if err != nil {
		log.Fatal(err)
	}

	var pngBuf bytes.Buffer
	_ = png.Encode(&pngBuf, image)

	var svgBuf bytes.Buffer
	err = gT.Render(graphV, graphviz.SVG, &svgBuf)

	err = createImages(pngBuf, svgBuf, outputFolder, outputName)

	if err != nil {
		log.Fatal(err)
	}

	// FINISH STATUS
	c := color.New(color.FgCyan, color.Bold)
	defer c.Println("Successfully created graph.")

	return nil
}

func createImages(pngBuf bytes.Buffer, svgBuf bytes.Buffer, outputFolder string, outputName string) error {
	svg, _ := os.Create(fmt.Sprintf("./%s/", outputFolder) + outputName + ".svg")
	svg.Write(svgBuf.Bytes())

	svg.Close()

	png, err := os.Create(fmt.Sprintf("./%s/", outputFolder) + outputName + ".png")
	if err != nil {
		log.Panic(err)
	}
	png.Write(pngBuf.Bytes())

	png.Close()

	return nil
}

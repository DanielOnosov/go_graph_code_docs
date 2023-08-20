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
)

func GenerateChart(nodes []structs.Node, edges []structs.Edge, outputFolder string, outputName string, params ...[]structs.Attr) {
	g := graph.New(graph.StringHash, graph.Directed())

	for _, node := range nodes {
		var attrs []func(properties *graph.VertexProperties)

		if len(params) >= 1 {
			for _, a := range params[0] {
				attrs = append(attrs, graph.VertexAttribute(a.Key, a.Value))
			}
		}

		for _, attr := range node.Attributes {
			attrs = append(attrs, graph.VertexAttribute(attr.Key, attr.Value))
		}

		attrs = append(attrs, graph.VertexAttribute("label", node.Label))

		_ = g.AddVertex(node.Id, attrs...)
	}

	for _, edge := range edges {
		var attrs []func(*graph.EdgeProperties)

		if len(params) >= 2 {
			for _, a := range params[1] {
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

	err = draw.DOT(g, file, draw.GraphAttribute("rankdir", "LR"), draw.GraphAttribute("bgcolor", "#111111"))

	if err != nil {
		fmt.Println(err.Error())
	}

	b, err := ioutil.ReadFile(fmt.Sprintf("./%s/", outputFolder) + outputName + ".gv")
	if err != nil {
		log.Fatal(err)
	}

	graphV, err := graphviz.ParseBytes(b)
	gT := graphviz.New()

	image, err := gT.RenderImage(graphV)
	if err != nil {
		log.Fatal(err)
	}

	var imageBuf bytes.Buffer
	err = png.Encode(&imageBuf, image)

	if err != nil {
		log.Fatal(err)
	}

	var svgBuf bytes.Buffer
	err = gT.Render(graphV, graphviz.SVG, &svgBuf)

	if err != nil {
		log.Fatal(err)
	}

	svg, err := os.Create(fmt.Sprintf("./%s/", outputFolder) + outputName + ".svg")
	svg.Write(svgBuf.Bytes())

	png, err := os.Create(fmt.Sprintf("./%s/", outputFolder) + outputName + ".png")
	if err != nil {
		log.Panic(err)
	}
	png.Write(imageBuf.Bytes())

	defer png.Close()

	// FINISH STATUS
	c := color.New(color.FgCyan, color.Bold)
	defer c.Println("Successfully created graph.")

}

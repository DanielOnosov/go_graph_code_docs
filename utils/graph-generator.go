package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/fatih/color"
	"github.com/goccy/go-graphviz"
	"graph-docs-golang/structs"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func GenerateChart(nodes []structs.Node, edges []structs.Edge, outputFolder string, outputName string) {
	g := graph.New(graph.StringHash, graph.Directed())

	for _, node := range nodes {
		var attrs []func(properties *graph.VertexProperties)

		for _, attr := range node.Attributes {
			attrs = append(attrs, graph.VertexAttribute(attr.Key, attr.Value))
		}

		attrs = append(attrs, graph.VertexAttribute("label", node.Label))

		_ = g.AddVertex(node.Id, attrs...)
	}

	for _, edge := range edges {
		var attrs []func(*graph.EdgeProperties)

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

	err = draw.DOT(g, file, draw.GraphAttribute("rankdir", "LR"))

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
		log.Panic(err)
	}

	fo, err := os.Create(fmt.Sprintf("./%s/", outputFolder) + outputName + ".png")

	defer fo.Close()
	c := color.New(color.FgCyan, color.Bold)
	defer c.Println("Successfully created graph.")

	if err != nil {
		log.Panic(err)
	}

	fw := bufio.NewWriter(fo)

	fw.Write(imageBuf.Bytes())
}
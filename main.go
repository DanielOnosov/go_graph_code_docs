package docParser

import (
	"github.com/DaksinWorld/go_graph_code_docs/structs"
	"github.com/DaksinWorld/go_graph_code_docs/utils"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Parser struct {
	Path         string
	OutputName   string
	OutputFolder string
	Title        string
	Theme        structs.Theme
}

func initParser(path string, outputFolder string, outputName string) *Parser {
	return &Parser{Path: path, OutputFolder: outputFolder, OutputName: outputName}
}

func Parse(path string, outputFolder string, outputName string) *Parser {
	return initParser(path, outputFolder, outputName)
}

func (p *Parser) Generate() error {
	entries, err := getChildren(p.Path, "")

	if err != nil {
		return err
	}

	var nodes []structs.Node
	var edges []structs.Edge

	for _, entry := range entries {
		content, err := ioutil.ReadFile(entry.Path)
		if err != nil {
			return err
		}

		contentStr := string(content)

		pattern := `//\s*DOC#\s*(.*?)\s*->\s*(.*?)\n`
		re := regexp.MustCompile(pattern)
		matches := re.FindAllStringSubmatch(contentStr, -1)

		// Get all lines, check if they exists
		if len(matches) >= 1 {
			for _, arrayWithNode := range matches {
				// Check if line has it own pattern
				if len(arrayWithNode) > 2 {
					newNodes, newEdges := CreateNodesAndEdges(p, nodes, edges, entry, arrayWithNode)

					nodes = newNodes
					edges = newEdges
				}
			}
		}
	}

	if len(nodes) <= 0 {
		c := color.New(color.BgRed, color.Bold, color.FgHiWhite)
		c.Println("Can't create graph, no data were collected")
	} else {
		err := utils.GenerateChart(nodes, edges, p.OutputFolder, p.OutputName, p.Title, p.Theme)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateNodesAndEdges(p *Parser, nodes []structs.Node, edges []structs.Edge, entry structs.Element, arrayWithNodeMatches []string) ([]structs.Node, []structs.Edge) {
	from := removePrefix(entry.Path, p.Path, "")
	relPathToSecondNode := removePrefix(arrayWithNodeMatches[2], "@", p.Path)
	to := removePrefix(relPathToSecondNode, p.Path, "")

	sourceIdx, targetIdx := findNodeByLabel(nodes, from, to)

	var node structs.Node

	var targetNode structs.Node

	// If Source Node already exists don't create a new one
	if sourceIdx != -1 {
		node = nodes[sourceIdx]
	} else {
		node = structs.Node{Id: from, Label: from}
	}

	// If Target Node already exists don't create a new one
	if targetIdx != -1 {
		targetNode = nodes[targetIdx]
	} else {
		targetNode = structs.Node{Id: to, Label: to}
	}

	var edge = structs.Edge{From: from, To: to}

	nodeAttrs, edgeAttrs := utils.OpenFileAndReturnNode(entry.Path)
	// Apply to root node and edge
	node.Attributes = nodeAttrs
	edge.Attributes = edgeAttrs

	// Apply to second node
	targetNodeAttrs, _ := utils.OpenFileAndReturnNode(relPathToSecondNode)
	targetNode.Attributes = targetNodeAttrs

	nodes = append(nodes, node, targetNode)
	edges = append(edges, edge)

	return nodes, edges
}

func (p *Parser) AddTitle(title string) {
	p.Title = title
}

func (p *Parser) AddTheme(theme structs.Theme) {
	p.Theme = theme
}

func removePrefix(s string, p string, v string) string {
	return strings.Replace(s, p, v, 1)
}

func getChildren(path string, name string) ([]structs.Element, error) {
	if name != "" {
		path = path + "/" + name
	}

	files, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		return nil, nil
	}

	var children []structs.Element

	for _, file := range files {
		if file.IsDir() {
			nestedChildren, err := getChildren(path, file.Name())

			if err != nil {
				log.Fatal(err)
			}

			children = append(children, nestedChildren...)
		} else {
			children = append(children, structs.Element{Path: path + "/" + file.Name()})
		}
	}

	return children, nil
}

func findNodeByLabel(nodes []structs.Node, from string, to string) (int, int) {
	for i := range nodes {
		if nodes[i].Label == from {
			return i, -1
		}
		if nodes[i].Label == to {
			return -1, i
		}
	}
	return -1, -1
}

package lib

import (
	"graph-docs-golang/structs"
	"graph-docs-golang/utils"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Parser struct {
	Path string
}

func InitParser(path string) *Parser {
	return &Parser{Path: path}
}

func Parse(path string) *Parser {
	return InitParser(path)
}

func (p *Parser) Generate() {
	entries, err := GetChildren(p.Path, "")

	if err != nil {
		log.Fatal(err)
	}

	var nodes []structs.Node
	var edges []structs.Edge

	id := 0

	for _, entry := range entries {
		content, err := ioutil.ReadFile(entry.Path)
		if err != nil {
			log.Fatal(err)
		}

		contentStr := string(content)

		pattern := `//\s*C##\s*(.*?):\s*(.*?)\n`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(contentStr)

		// Desc patterns
		descPatter := `//\s*C##\s*(.*?)=\s*(.*?):\s*(.*?)(?:\n|$)`
		descRe := regexp.MustCompile(descPatter)
		descMatches := descRe.FindAllStringSubmatch(contentStr, -1)

		if len(matches) > 2 {
			from := RemovePrefix(entry.Path, p.Path, "")
			id += 1

			var node = structs.Node{Id: id, Label: from}

			to := RemovePrefix(RemovePrefix(matches[2], "@", p.Path), p.Path, "")
			id += 1

			var secondNode = structs.Node{Id: id, Label: to}

			var edge = structs.Edge{From: id - 1, To: id}
			var attrs []structs.Attr

			var lastAttrType = ""

			for _, el := range descMatches {
				m := descRe.FindStringSubmatch(el[0])

				fKey := strings.TrimSpace(m[1]) // type of attr
				sKey := strings.TrimSpace(m[2]) // key
				tKey := strings.TrimSpace(m[3]) // value

				if lastAttrType != "" {
					if lastAttrType != fKey {
						attrs = attrs[:0]
					}
				}

				attrs = append(attrs, structs.Attr{
					Key:   sKey,
					Value: tKey,
				})

				if fKey == "Node" {
					node.Attributes = append(node.Attributes, attrs...)
					lastAttrType = "Node"
				} else {
					edge.Attributes = append(edge.Attributes, attrs...)
					lastAttrType = "Edge"
				}
			}

			nodes = append(nodes, node, secondNode)
			edges = append(edges, edge)
		}
	}
	//
	//fmt.Println(edges)
	//fmt.Println(nodes)

	utils.GenerateChart(nodes, edges)
}

func RemovePrefix(s string, p string, v string) string {
	return strings.Replace(s, p, v, 1)
}

func GetChildren(path string, name string) ([]structs.Element, error) {
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
			nestedChildren, err := GetChildren(path, file.Name())

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

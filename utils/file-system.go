package utils

import (
	"github.com/DaksinWorld/go_graph_code_docs/structs"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func CreateOutputFolder(path string) {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := os.RemoveAll(path)

		if err != nil {
			log.Fatal(err)
		}

		err = os.Mkdir(path, 0777)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func OpenFileAndReturnNode(path string) ([]structs.Attr, []structs.Attr) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	contentStr := string(content)

	descPatter := `//\s*DOC##\s*(.*?)=\s*(.*?):\s*(.*?)(?:\n|$)`
	descRe := regexp.MustCompile(descPatter)
	descMatches := descRe.FindAllStringSubmatch(contentStr, -1)

	var edgeAttrs []structs.Attr
	var nodeAttrs []structs.Attr

	for _, el := range descMatches {
		m := descRe.FindStringSubmatch(el[0])

		fKey := strings.TrimSpace(m[1]) // type of attr
		sKey := strings.TrimSpace(m[2]) // key
		tKey := strings.TrimSpace(m[3]) // value

		if fKey == "Node" {
			nodeAttrs = append(nodeAttrs, structs.Attr{
				Key:   sKey,
				Value: tKey,
			})
		} else {
			edgeAttrs = append(edgeAttrs, structs.Attr{
				Key:   sKey,
				Value: tKey,
			})
		}
	}

	return nodeAttrs, edgeAttrs
}

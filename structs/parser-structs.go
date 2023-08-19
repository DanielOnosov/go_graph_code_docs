package structs

import "os"

type Element struct {
	Path string
	Data os.DirEntry
}

type Node struct {
	Attributes []Attr
	Id         int
	Label      string
}

type Attr struct {
	Key   string
	Value string
}

type Edge struct {
	Attributes []Attr
	From       int
	To         int
}

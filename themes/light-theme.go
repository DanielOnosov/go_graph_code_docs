package themes

import "github.com/DaksinWorld/go_graph_code_docs/structs"

var nodeTemplate = []structs.Attr{
	{Key: "height", Value: "2"},
	{Key: "shape", Value: "box"},
	{Key: "style", Value: "rounded,filled"},
	{Key: "fontcolor", Value: "#1F1F1F"},
	{Key: "fillcolor", Value: "#dce8e4"},
	{Key: "fontname", Value: "Montserrat"},
	{Key: "fontsize", Value: "10"},
	{Key: "color", Value: "black"},
	{Key: "width", Value: "2"},
}
var edgeTemplate = []structs.Attr{
	{Key: "color", Value: "#1F1F1F"},
	{Key: "fontcolor", Value: "#1F1F1F"},
	{Key: "fontname", Value: "Montserrat"},
	{Key: "fontweight", Value: "700"},
	{Key: "fontsize", Value: "10"},
}

var LightTheme = structs.Theme{Name: "LightTheme", CollOfAttrs: [][]structs.Attr{nodeTemplate, edgeTemplate}}

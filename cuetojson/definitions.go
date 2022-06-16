package cuetojson

import (
	"dagviz/dag"
	"os"
	"regexp"
)

func LinkDefinitions(infos []CueInfos, root *dag.Root) {
	definitionsNode := root.AttachNode("definitions")

	for _, program := range infos {
		for _, file := range program.Files {
			content, err := os.ReadFile(file)

			if err != nil {
				continue
			}

			definitions := parseDefinitions(string(content))
			findDefinitions(definitions, definitionsNode)
		}
	}
}

func parseDefinitions(content string) []string {
	regex := regexp.MustCompile("#[^ ]+")
	array := regex.FindAllString(content, -1)

	return array
}

func findDefinitions(definitions []string, node *dag.Node) {
	for _, definition := range definitions {
		def := &dag.Node{Value: definition}
		node.LinksTo(def)
	}
}

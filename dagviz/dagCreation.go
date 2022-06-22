package dagviz

import (
	"dagviz/dag"
	"fmt"
)

func LinkDefinitions(infos []CueInfos, root *CueRoot) {
	definitionsNode := root.AttachNode(NodeDefinition{
		name: "definitions",
		file: "",
		def:  "",
	})

	for _, program := range infos {
		for _, file := range program.Files {
			content, err := getCueContent(file)

			if err != nil {
				continue
			}

			dependencies := program.getDependencies()
			dependencies = append(dependencies, file)
			buildFiles := sortDependencies(dependencies)
			p, err := getPackage(file)

			if err != nil {
				fmt.Println(err)
				continue
			}

			definitions := parseDefinitions(content, p)
			addDefinitionsToDag(definitions, buildFiles, program.Root, definitionsNode)
		}
	}
}

func addDefinitionsToDag(definitions []DefinitionData, buildFiles map[string][]string, root string, node *dag.Node) {
	for _, definition := range definitions {
		getDefinitions(node, buildFiles, root, definition)
	}
}

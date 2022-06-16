package cuetojson

import (
	"dagviz/dag"
	"fmt"
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
			addDefinitionsToDag(definitions, definitionsNode)
		}
	}
}

func addDefinitionsToDag(definitions []string, node *dag.Node) {
	for _, definition := range definitions {
		def := &dag.Node{Value: definition}
		node.LinksTo(def)
	}
}

func parseDefinitions(content string) []string {
	regex := regexp.MustCompile("#[^ ]+")
	array := regex.FindAllString(content, -1)

	return array
}

func defineNeedle(file string, needle string) (bool, Definition) {
	byteFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return false, Definition{}
	}
	stringFile := string(byteFile)
	reg := regexp.MustCompile(fmt.Sprintf("\n%s: {", needle))
	//regDef := regexp.MustCompile(fmt.Sprintf("\n%s: .+[\n}]/gms", needle))
	if reg.MatchString(stringFile) {
		return true, Definition{file, ""}
	}
	return false, Definition{}
}

func findDefinition(files []string, needle string) (Definition, error) {
	for _, file := range files {
		if boole, def := defineNeedle(file, needle); boole {
			return def, nil
		}
	}
	return Definition{}, fmt.Errorf("could not find definition for %s", needle)
}

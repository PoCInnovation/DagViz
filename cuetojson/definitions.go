package cuetojson

import (
	"dagviz/dag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func LinkDefinitions(infos []CueInfos, root *dag.Root) {
	definitionsNode := root.AttachNode("definitions")

	for _, program := range infos {
		for _, file := range program.Files {
			content, err := os.ReadFile(file)

			if err != nil {
				continue
			}

			deps := program.getDependencies()
			deps = append(deps, file)
			definitions := parseDefinitions(string(content))
			addDefinitionsToDag(definitions, deps, program.Root, definitionsNode)
		}
	}
}

type NodeDefinition struct {
	name string
	file string
	def  string
}

func addDefinitionsToDag(definitions []string, buildFiles []string, root string, node *dag.Node) {
	for _, definition := range definitions {
		defNode := &dag.Node{}
		node.LinksTo(defNode)

		data, err := findDefinition(buildFiles, definition)
		if err != nil {
			continue
		}

		defNode.Value = NodeDefinition{
			name: definition,
			file: data.file,
			def:  "definition",
		}
		go addDefinitionsToDag(parseDefinitions(data.def), buildFiles, root, defNode)
	}
}

func parseDefinitions(content string) []string {
	var definitions []string

	regex := regexp.MustCompile("#[^ ,\n]+")
	array := regex.FindAllString(content, -1)

	for _, def := range array {
		if !strings.HasSuffix(def, ":") {
			definitions = append(definitions, def)
		}
	}

	return definitions
}

func findDefinition(files []string, needle string) (Definition, error) {
	for _, file := range files {
		if boole, def := defineNeedle(file, needle); boole {
			return def, nil
		}
	}
	return Definition{}, fmt.Errorf("could not find definition for %s", needle)
}

func defineNeedle(file string, needle string) (bool, Definition) {
	byteFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return false, Definition{}
	}
	stringFile := string(byteFile)
	reg := regexp.MustCompile("\n" + needle + ": ")
	if reg.MatchString(stringFile) {
		def, err := extractDefinition(stringFile, needle)
		if err != nil {
			fmt.Println(err)
			return false, Definition{}
		}
		return true, Definition{file, def}
	}
	return false, Definition{}
}

func extractDefinition(file string, needle string) (string, error) {
	regDef := regexp.MustCompile("(?s)\n" + needle + ": .+")
	def := regDef.FindString(file)
	for i, _ := range def {
		if i > len(def)-1 {
			return "", fmt.Errorf("could not find definition for %s", needle)
		}

		if def[i] == '\n' && def[i+1] == '}' {
			return def[:i+1], nil
		}
	}
	return "", fmt.Errorf("could not find definition for %s", needle)
}

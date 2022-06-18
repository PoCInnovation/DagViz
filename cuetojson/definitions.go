package cuetojson

import (
	"dagviz/dag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func LinkDefinitions(infos []CueInfos, root *CueRoot) {
	definitionsNode := root.AttachNode(NodeDefinition{
		name: "definitions",
		file: "",
		def:  "",
	})

	for _, program := range infos {
		for _, file := range program.Files {
			content, err := os.ReadFile(file)

			if err != nil {
				continue
			}

			deps := program.getDependencies()
			deps = append(deps, file)
			definitions := parseDefinitions(string(content))

			fmt.Println("def of: ", definitionsNode.Value.(NodeDefinition).name)
			addDefinitionsToDag(definitions, deps, program.Root, definitionsNode)
			time.Sleep(time.Millisecond * 1000)
		}
	}
}

type NodeDefinition struct {
	name string
	file string
	def  string
}

func getDefinitions(node *dag.Node, buildFiles []string, root string, definition string) {
	defNode := &dag.Node{}

	node.LinksTo(defNode)

	data, err := findDefinition(buildFiles, definition)
	if err != nil {
		return
	}

	defNode.Value = NodeDefinition{
		name: definition,
		file: strings.Replace(data.file, root, "", -1),
		def:  "definition",
	}
	fmt.Println("def in: ", data.def)
	addDefinitionsToDag(parseDefinitions(data.def), buildFiles, root, defNode)
}

var debug int

func addDefinitionsToDag(definitions []string, buildFiles []string, root string, node *dag.Node) {
	for _, def := range definitions {
		fmt.Println("adding definition:", def)
	}
	fmt.Print("\n\n")

	for _, definition := range definitions {
		getDefinitions(node, buildFiles, root, definition)
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

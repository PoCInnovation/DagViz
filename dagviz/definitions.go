package dagviz

import (
	"dagviz/dag"
	"fmt"
	"golang.org/x/exp/slices"
	"os"
	"regexp"
	"strings"
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
			definitions := parseDefinitions(string(content), p)

			addDefinitionsToDag(definitions, buildFiles, program.Root, definitionsNode)
		}
	}
}

func getPackage(file string) (string, error) {
	byteFile, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	fileString := string(byteFile)
	reg := regexp.MustCompile("package ([a-zA-Z0-9]+)")
	p := reg.FindStringSubmatch(fileString)

	return p[1], nil
}

func sortDependencies(dependencies []string) map[string][]string {
	var sortedDependencies = make(map[string][]string)

	for _, d := range dependencies {
		pack, err := getPackage(d)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if _, ok := sortedDependencies[pack]; !ok {
			sortedDependencies[pack] = []string{}
		}
		sortedDependencies[pack] = append(sortedDependencies[pack], d)
	}
	return sortedDependencies
}

func getDefinitions(node *dag.Node, buildFiles map[string][]string, root string, definition DefinitionData) {
	defNode := &dag.Node{}

	node.LinksTo(defNode)

	data, err := findDefinition(buildFiles[definition.pack], definition.defName)
	if err != nil {
		return
	}

	defNode.Value = NodeDefinition{
		name: definition.defName,
		file: strings.Replace(data.file, root, "", -1),
		def:  data.def,
	}
	addDefinitionsToDag(parseDefinitions(data.def, definition.pack), buildFiles, root, defNode)
}

func addDefinitionsToDag(definitions []DefinitionData, buildFiles map[string][]string, root string, node *dag.Node) {
	for _, definition := range definitions {
		getDefinitions(node, buildFiles, root, definition)
	}
}

func parseDefinitions(content string, pack string) []DefinitionData {
	var definitions []DefinitionData

	regex := regexp.MustCompile("([a-zA-Z.]*)(#[^ ,\n]+)")
	array := regex.FindAllStringSubmatch(content, -1)

	for _, def := range array {
		if len(def[1]) == 0 {
			if !strings.HasSuffix(def[2], ":") {
				if slices.Contains(definitions, DefinitionData{def[2], pack}) {
					continue
				}
				definitions = append(definitions, DefinitionData{def[2], pack})
			}
		} else {
			if slices.Contains(definitions, DefinitionData{def[2], def[1][:len(def[1])-1]}) {
				continue
			}
			definitions = append(definitions, DefinitionData{def[2], def[1][:len(def[1])-1]})
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
		return true, Definition{file, def + "}"}
	}
	return false, Definition{}
}

func extractDefinition(file string, needle string) (string, error) {
	regDef := regexp.MustCompile("(?s)\n" + needle + ": .+")
	def := regDef.FindString(file)
	for i := range def {
		if i > len(def)-1 {
			return "", fmt.Errorf("could not find definition for %s", needle)
		}

		if def[i] == '\n' && def[i+1] == '}' {
			return def[:i+1], nil
		}
	}
	return "", fmt.Errorf("could not find definition for %s", needle)
}

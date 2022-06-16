package main

import (
	"fmt"
	"os"
	"regexp"
)

type definition struct {
	file string
	def  string
}

func defineNeedle(file string, needle string) (bool, definition) {
	byteFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return false, definition{}
	}
	stringFile := string(byteFile)
	reg := regexp.MustCompile(fmt.Sprintf("\n%s: {", needle))
	//regDef := regexp.MustCompile(fmt.Sprintf("\n%s: .+[\n}]/gms", needle))
	if reg.MatchString(stringFile) {
		return true, definition{file, ""}
	}
	return false, definition{}
}

func findDefinition(files []string, needle string) (definition, error) {
	for _, file := range files {
		if boole, def := defineNeedle(file, needle); boole {
			return def, nil
		}
	}
	return definition{}, fmt.Errorf("could not find definition for %s", needle)
}

package dagviz

import (
	"fmt"
	"os"
	"regexp"
)

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

package dagviz

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/build"
	"os"
	"regexp"
)

func ExtractInfos(programs []CueProgram) []CueInfos {
	var infos []CueInfos

	for _, program := range programs {
		var importList []CueInfos

		info := CueInfos{}
		info.Tags = program.Instance.AllTags
		info.Root = program.Instance.Root
		info.Module = program.Instance.Module
		info.Package = program.Instance.PkgName
		info.Dependencies = program.Instance.Deps
		info.Directory = program.Instance.Dir
		info.Files = extractFiles(program.Instance.Files)
		info.BuildFiles = extractBuildFiles(program.Instance.BuildFiles)
		info.InvalidFiles = extractBuildFiles(program.Instance.InvalidFiles)
		info.IgnoredFiles = extractBuildFiles(program.Instance.IgnoredFiles)
		info.OrphanedFiles = extractBuildFiles(program.Instance.OrphanedFiles)

		for _, imports := range program.Instance.Imports {
			program := CueProgram{imports, nil}
			sub := ExtractInfos([]CueProgram{program})
			importList = append(infos, sub...)
		}
		info.Imports = importList

		infos = append(infos, info)
	}
	return infos
}

func extractFiles(files []*ast.File) []string {
	var array []string

	if files == nil {
		return make([]string, 0)
	}
	for _, file := range files {
		array = append(array, file.Filename)
	}

	return array
}

func extractBuildFiles(files []*build.File) []string {
	var array []string

	if files == nil {
		return make([]string, 0)
	}
	for _, file := range files {
		array = append(array, file.Filename)
	}

	return array
}

func getCueContent(file string) (string, error) {
	regex := regexp.MustCompile("(?s)//.*?\\n|/\\\\*.*?\\\\*/")
	content, err := os.ReadFile(file)

	if err != nil {
		return "", err
	}

	uncomment := regex.ReplaceAll(content, nil)
	return string(uncomment), nil
}

func (infos CueInfos) getDependencies() []string {
	var array []string

	for _, sub := range infos.Imports {
		array = append(array, sub.Files...)
		imports := sub.getDependencies()
		array = append(array, imports...)
	}
	return array
}

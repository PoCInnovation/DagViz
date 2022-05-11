package cuetojson

import (
	"cuelang.org/go/cue/ast"
)

func extractFiles(files []*ast.File) []string {
	var array []string

	for _, file := range files {
		array = append(array, file.Filename)
	}

	return array
}

func ExtractInfos(programs []CueProgram) []CueInfos {
	var infos []CueInfos

	for _, program := range programs {
		info := CueInfos{}
		info.Values = program.Values
		info.Tags = program.Instance.AllTags
		info.Root = program.Instance.Root
		info.Module = program.Instance.Module
		info.Package = program.Instance.PkgName
		info.Dependencies = program.Instance.Deps
		info.Files = extractFiles(program.Instance.Files)
		info.Directory = program.Instance.Dir

		infos = append(infos, info)
	}
	return infos
}

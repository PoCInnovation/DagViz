package cuetojson

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/build"
)

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

package cuetojson

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/build"
)

type CueProgram struct {
	Instance *build.Instance
	Values   *cue.Value
}

type CueInfos struct {
	Tags          []string
	Root          string
	Module        string
	Package       string
	Dependencies  []string
	Directory     string
	Files         []string
	BuildFiles    []string
	InvalidFiles  []string
	IgnoredFiles  []string
	OrphanedFiles []string
	Imports       []CueInfos
}

type Definition struct {
	file string
	def  string
}

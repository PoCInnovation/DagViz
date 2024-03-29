package cueToDag

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/build"
	"github.com/PoCInnovation/DagViz/dag"
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

type CueRoot struct {
	dag.Root
}

type Definition struct {
	file string
	def  string
}

type NodeDefinition struct {
	name string
	file string
	def  string
}

type DefinitionData struct {
	defName string
	pack    string
}

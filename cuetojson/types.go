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
	Tags         []string `json:"tags"`
	Root         string   `json:"root"`
	Module       string   `json:"module"`
	Package      string   `json:"package"`
	Dependencies []string `json:"dependencies"`
	Directory    string   `json:"directory"`
	Files        []string `json:"files"`
	BuildFiles   []string `json:"build_files"`
}

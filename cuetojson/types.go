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
	Values *cue.Value `json:"values"`
}

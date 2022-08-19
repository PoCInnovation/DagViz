package cueToDag

import (
	"testing"

	"github.com/PoCInnovation/DagViz/dag"

	"cuelang.org/go/cue/cuecontext"
	"github.com/stretchr/testify/assert"
)

func TestCreateCueDag(t *testing.T) {
	CueDir := "../samples/package_test"
	context := cuecontext.New()
	programs, err := LoadFile(context, CueDir, nil)
	if err != nil {
		t.Fail()
	}

	infos := ExtractInfos(programs)
	root := CreateCueDag("dag")
	LinkDefinitions(infos, &root)
	expected := []*dag.Node{
		{
			Value: NodeDefinition{
				name: "#Image",
				file: "/cue.mod/pkg/universe.dagger.io/docker/image.cue",
				def:  "#Image: {\n\t// Root filesystem of the image.\n\trootfs: dagger.#FS\n\n\t// Image config\n\tconfig: core.#ImageConfig\n}",
			},
			Links: []*dag.Node{
				{
					Value: NodeDefinition{
						name: "#FS",
						file: "/cue.mod/pkg/dagger.io/dagger/types.cue",
						def:  "#FS: {\n\t$dagger: fs: _id: string | null\n}",
					},
				},
				{
					Value: NodeDefinition{
						name: "#ImageConfig",
						file: "/cue.mod/pkg/dagger.io/dagger/core/image.cue",
						def:  "#ImageConfig: {\n\tuser?: string\n\texpose?: [string]: {}\n\tenv?: [string]: string\n\tentrypoint?: [...string]\n\tcmd?: [...string]\n\tvolume?: [string]: {}\n\tworkdir?: string\n\tlabel?: [string]: string\n\tstopsignal?:  string\n\thealthcheck?: #HealthCheck\n\targsescaped?: bool\n\tonbuild?: [...string]\n\tstoptimeout?: int\n\tshell?: [...string]\n}",
					},
					Links: []*dag.Node{
						{
							Value: NodeDefinition{
								name: "#HealthCheck",
								file: "/cue.mod/pkg/dagger.io/dagger/core/image.cue",
								def:  "#HealthCheck: {\n\ttest?: [...string]\n\tinterval?:    int\n\ttimeout?:     int\n\tstartperiod?: int\n\tretries?:     int\n}",
							},
						},
					},
				},
			},
		},
	}
	assert.Equal(t, expected, root.Members[0].Links, "got %v but expected %v", root.Members[0], expected)
}

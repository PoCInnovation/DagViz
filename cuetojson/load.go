package cuetojson

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
	"fmt"
)

func LoadFile(context *cue.Context, path string) ([]CueProgram, error) {
	singleton := []string{"dagger/cue.mod/pkg/universe.dagger.io/git/git.cue"}
	config := load.Config{
		Dir:        ".",
		Package:    "*",
		ModuleRoot: ".",
		Module:     "root",
	}
	instances := load.Instances(singleton, &config)
	var programs []CueProgram

	fmt.Printf("%d %d\n", len(instances), len(instances[0].Files))
	if len(instances) == 0 {
		return nil, fmt.Errorf("Error: Cannot load %s\n", path)
	}

	for _, instance := range instances {
		if instance.Err != nil {
			fmt.Printf("Error: Cannot load %s\n%s\n", path, instance.Err.Error())
			return nil, instance.Err
		}
		value := context.BuildInstance(instance)
		if value.Err() != nil {
			fmt.Printf("Error: Cannot build %s\n%s\n", path, value.Err().Error())
			return nil, value.Err()
		}

		validation := value.Validate()
		if validation != nil {
			fmt.Printf("Error: Cannot validate %s\n%s\n", path, validation.Error())
			return nil, value.Err()
		}
		programs = append(programs, CueProgram{instance, &value})
	}

	return programs, nil
}

package cuetojson

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/build"
	"cuelang.org/go/cue/load"
	"fmt"
)

func LoadFile(context *cue.Context, path string) (*cue.Value, error) {
	var instance *build.Instance
	var value cue.Value
	var validation error
	singleton := []string{path}
	instances := load.Instances(singleton, nil)

	if len(instances) == 0 {
		return nil, fmt.Errorf("Error: Cannot load %s\n", path)
	}

	instance = instances[0]
	if instance.Err != nil {
		fmt.Printf("Error: Cannot load %s\n%s\n", path, instance.Err.Error())
		return nil, instance.Err
	}

	value = context.BuildInstance(instance)
	if value.Err() != nil {
		fmt.Printf("Error: Cannot build %s\n%s\n", path, value.Err().Error())
		return nil, value.Err()
	}

	validation = value.Validate()
	if validation != nil {
		fmt.Printf("Error: Cannot validate %s\n%s\n", path, validation.Error())
		return nil, value.Err()
	}

	return &value, nil
}

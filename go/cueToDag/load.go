package cueToDag

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/load"
)

func LoadFile(context *cue.Context, p string, overlay map[string]fs.FS) ([]CueProgram, error) {
	config := &load.Config{
		Dir:     p,
		Overlay: map[string]load.Source{},
	}

	for mnt, f := range overlay {
		err := fs.WalkDir(f, ".", func(p string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !entry.Type().IsRegular() || filepath.Ext(entry.Name()) != ".cue" {
				return nil
			}

			contents, err := fs.ReadFile(f, p)
			if err != nil {
				return fmt.Errorf("%s: %w", p, err)
			}

			overlayPath := path.Join(config.Dir, mnt, p)
			config.Overlay[overlayPath] = load.FromBytes(contents)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return loadCueFiles(context, config, p)
}

func loadCueFiles(context *cue.Context, config *load.Config, path string) ([]CueProgram, error) {
	instances := load.Instances([]string{"."}, config)
	var programs []CueProgram

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

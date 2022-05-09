package lib

import (
	s "dag/structure"
)

func CreateDag(name string) *s.Root {
	return &s.Root{Name: name}
}

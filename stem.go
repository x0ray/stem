package stem

import (
	"io"
	"strings"
)

const (
	sep = "."
)

type Stem struct {
	name  string
	syms  map[string]any
	owner *Stem
}

func NewStem(name string, owner *Stem) *Stem {
	stem := new(Stem)
	stem.name = name
	stem.syms = make(map[string]any)
	stem.owner = owner
	return stem
}

func verify(path string) error {
	// TODO verify form ccc or ccc.ccc.ccc. ... ccc
	return nil
}

func (s *Stem) Set(path string, value any) error {
	if err := verify(path); err != nil {
		return err
	}
	s.syms[s.name] = value
	parts := strings.Split(path, sep)
	// TODO
	ns := NewStem(part, s)
	ns.Set()
	// TODO
}

func (s *Stem) Get(path string) any {
	// TODO
	return nil
}

func (s *Stem) Exist(path string) bool {
	// TODO
	return false
}

func (s *Stem) Delete(path string) error {
	// TODO
	return nil
}

func (s *Stem) Print(o io.Writer) error {
	// TODO
	return nil
}

func (s *Stem) MarshalJSON(path string) []byte {
	// TODO
	return nil
}

func (s *Stem) UnmarshalJSON(in []byte) error {
	// TODO
	return nil
}

func (s *Stem) Increment(path string, val any) {
	// TODO
}

func (s *Stem) Next(path string) any {
	// TODO
	return nil
}

func (s *Stem) First(path string) any {
	// TODO
	return nil
}

func (s *Stem) Last(path string) any {
	// TODO
	return nil
}

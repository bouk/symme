package symme

import (
	"debug/gosym"
	"github.com/bouk/symme/objfile"
	"github.com/kardianos/osext"
)

func Table() (*gosym.Table, error) {
	path, err := osext.Executable()
	if err != nil {
		return nil, err
	}
	f, err := objfile.Open(path)
	if err != nil {
		return nil, err
	}
	return f.PCLineTable()
}

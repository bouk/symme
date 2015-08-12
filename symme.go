package symme

import (
	"debug/gosym"
	"github.com/bouk/symme/objfile"
	"github.com/kardianos/osext"
	"sync"
)

var (
	f        *objfile.File
	initOnce sync.Once
)

func initobjfile() {
	initOnce.Do(func() {
		var err error
		path, err := osext.Executable()
		if err != nil {
			panic(err)
		}
		f, err = objfile.Open(path)
		if err != nil {
			panic(err)
		}
	})
}

func Table() (*gosym.Table, error) {
	initobjfile()
	return f.PCLineTable()
}

func Symbols() ([]objfile.Sym, error) {
	initobjfile()
	return f.Symbols()
}

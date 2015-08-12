package symme

import (
	"testing"
)

func TestTable(t *testing.T) {
	_, err := Table()
	if err != nil {
		t.Error(err)
	}
}

func TestSymbols(t *testing.T) {
	symbols, err := Symbols()
	if err != nil {
		t.Error(err)
	}
}

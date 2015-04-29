package symme

import (
	"testing"
)

func TestSymme(t *testing.T) {
	_, err := Table()
	if err != nil {
		t.Error(err)
	}
}

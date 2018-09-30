package somepkg

import (
	"testing"
)

func TestTilesKinds(t *testing.T) {
	arr := []Kind{Tree, Dirt, Water}
	for _, v := range arr {
		t.Logf("%T %#v \n", v, v)
	}
}

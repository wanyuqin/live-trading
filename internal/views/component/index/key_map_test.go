package index

import "testing"

func TestIndexKeyMap_Help(t *testing.T) {
	print(newIndexKeyMap().Help())
}

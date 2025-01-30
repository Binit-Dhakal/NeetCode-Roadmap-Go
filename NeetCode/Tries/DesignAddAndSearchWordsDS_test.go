package tries_test

import (
	"testing"

	tries "github.com/Binit-Dhakal/LeetCode-Go/Tries"
)

func TestDesignAddAndSearchDS(t *testing.T) {
	obj := tries.ConstructWD()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")

	if obj.Search("pad") != false {
		t.Error("pad shouldnot have been found")
	}

	if obj.Search("bad") != true {
		t.Error("bad is not found")
	}

	if obj.Search(".ad") != true {
		t.Error(".ad should result in true but is false")
	}

	if obj.Search("b..") != true {
		t.Error("b.. should result in true, but is false")
	}
}

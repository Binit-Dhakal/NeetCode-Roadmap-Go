package tries_test

import (
	"testing"

	tries "github.com/Binit-Dhakal/LeetCode-Go/Tries"
)

func TestImplementTries(t *testing.T) {
	try := tries.Constructor()
	try.Insert("apple")
	if try.Search("apple") != true {
		t.Errorf("just inserted apple not found")
	}

	if try.Search("app") != false {
		t.Errorf("app found although it was not inserted")
	}

	if try.StartsWith("app") != true {
		t.Errorf("app prefix not found although apple is there")
	}

	try.Insert("app")
	if try.Search("app") != true {
		t.Errorf("Inserted app not found")
	}
}

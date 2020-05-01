package supervillains

import "testing"

func TestAll(t *testing.T) {
	supervillains := All()

	if len(supervillains) == 0 {
		t.Errorf("Expect list to not be empty")
	}
}

func TestRandom(t *testing.T) {
	randomSupervillain := Random()

	if randomSupervillain == "" {
		t.Errorf("Expect random supervillain to exist")
	}
}

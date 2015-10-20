package badge

import (
	"os"
	"testing"
)

func TestDirExistsReturnNoError(t *testing.T) {
	cases := []struct {
		assert string
		expect bool
	}{
		{os.Getenv("HOME"), true},
		{"does/not/exist", false},
		{"./", true},
	}

	for _, c := range cases {
		got := dirExists(c.assert)

		if got != c.expect {
			t.Errorf("Error asserting that %t is equals to %t", got, c.expect)
		}
	}
}

package slices_test

import (
	"testing"

	"github.com/lsymds/go-utils/slices"
)

func TestMin(t *testing.T) {
	tcs := []struct {
		name string
		val  []int
		res  int
	}{
		{name: "integers", val: []int{10, 15, 1, 8, 3}, res: 1},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if res, err := slices.Min(tc.val); err != nil {
				t.Fatal("did not expect error to be returned from slices.Min")
			} else if res != tc.res {
				t.Fatalf("expected min to return %d, got: %d", tc.res, res)
			}
		})
	}
}

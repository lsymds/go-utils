package slices_test

import (
	"testing"

	"github.com/lsymds/go-utils/slices"
)

func TestMap(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		type user struct {
			id   int
			name string
		}

		res, err := slices.Map(
			[]user{
				{id: 1, name: "Foo"},
				{id: 2, name: "Bar"},
				{id: 3, name: "Baz"},
			},
			func(user *user) int {
				return user.id
			},
		)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		} else if res[0] != 1 {
			t.Fatalf("first mapped element was not 1, got: %d", res[0])
		}
	})
}

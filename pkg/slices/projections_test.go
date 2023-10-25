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

func TestFlatMap(t *testing.T) {
	type orderItem struct {
		value int
	}

	type order struct {
		items []orderItem
	}

	res, err := slices.FlatMap(
		[]order{
			{items: []orderItem{{value: 5}}},
			{items: []orderItem{{value: 10}, {value: 15}}},
			{items: []orderItem{{value: 85}}},
			{items: []orderItem{{value: 32}}},
		},
		func(el *order) []orderItem {
			return el.items
		},
	)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	} else if len(res) != 5 {
		t.Fatalf("count of flatmapped items was not 5, got: %d", len(res))
	}
}

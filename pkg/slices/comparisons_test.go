package slices_test

import (
	"testing"

	"github.com/lsymds/go-utils/slices"
)

func TestMin(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		if res, err := slices.Min([]int{10, 15, 1, 8, 3}); err != nil {
			t.Fatal("did not expect error to be returned from slices.Min")
		} else if res != 1 {
			t.Fatalf("expected min to return 1, got: %d", res)
		}
	})

	t.Run("floats", func(t *testing.T) {
		if res, err := slices.Min([]float32{8.384, 1100.394, 8.383, 7.4812, 7.4810849}); err != nil {
			t.Fatal("did not expect error to be returned from slices.Min")
		} else if res != 7.4810849 {
			t.Fatalf("expected min to return 7.4810849, got: %f", res)
		}
	})

	t.Run("derived types", func(t *testing.T) {
		type d int64

		if res, err := slices.Min([]d{8, 123981289389123, 588585858585, 9393939393939, 32020, 959049050909340922}); err != nil {
			t.Fatal("did not expect error to be returned from slices.Min")
		} else if res != 8 {
			t.Fatalf("expected min to return 8, got: %d", res)
		}
	})
}

func TestMinBy(t *testing.T) {
	type s struct {
		id   int
		name string
	}

	res, err := slices.MinBy(
		[]s{
			{id: 350, name: "Baz"},
			{id: 8384, name: "Foo"},
			{id: 90389, name: "Bar"},
		},
		func(el *s) int {
			return el.id
		},
	)
	if err != nil {
		t.Fatal("did not expect error to be returned from slices.Min")
	} else if res.id != 350 {
		t.Fatalf("expected minby to return id of 348, got: %d", res.id)
	} else if res.name != "Baz" {
		t.Fatalf("expected minby to return name of Baz, got: %s", res.name)
	}
}

func TestMax(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		if res, err := slices.Max([]int{10, 15, 1, 8, 3}); err != nil {
			t.Fatal("did not expect error to be returned from slices.Min")
		} else if res != 15 {
			t.Fatalf("expected max to return 15, got: %d", res)
		}
	})

	t.Run("floats", func(t *testing.T) {
		if res, err := slices.Max([]float32{8.384, 1100.394, 8.383, 7.4812, 7.4810849}); err != nil {
			t.Fatal("did not expect error to be returned from slices.Min")
		} else if res != 1100.394 {
			t.Fatalf("expected min to return 1100.394, got: %f", res)
		}
	})

	t.Run("derived types", func(t *testing.T) {
		type d int64

		if res, err := slices.Max([]d{8, 123981289389123, 588585858585, 9393939393939, 32020, 959049050909340922}); err != nil {
			t.Fatal("did not expect error to be returned from slices.Min")
		} else if res != 959049050909340922 {
			t.Fatalf("expected min to return 959049050909340922, got: %d", res)
		}
	})
}

func TestMaxBy(t *testing.T) {
	type s struct {
		id   int
		name string
	}

	res, err := slices.MaxBy(
		[]s{
			{id: 8384, name: "Foo"},
			{id: 90389, name: "Bar"},
			{id: 350, name: "Baz"},
		},
		func(el *s) int {
			return el.id
		},
	)
	if err != nil {
		t.Fatal("did not expect error to be returned from slices.Min")
	} else if res.id != 90389 {
		t.Fatalf("expected MaxBy to return id of 90389, got: %d", res.id)
	} else if res.name != "Bar" {
		t.Fatalf("expected MaxBy to return name of Bar, got: %s", res.name)
	}
}

package argsort

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestArgsort(t *testing.T) {
	orig := []string{"z", "a", "c", "b"}
	copy := append([]string{}, orig...)

	indices := Sort(sort.StringSlice(orig))
	if !reflect.DeepEqual(orig, copy) {
		t.Fatalf("argsort should not mutate original slice: %q != %q", orig, copy)
	}

	sorted := copy
	sort.Sort(sort.StringSlice(sorted))

	for i := range orig {
		if got, want := orig[indices[i]], sorted[i]; got != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	}
}

func ExampleSort() {
	in := []string{"z", "a", "c", "b"}
	order := Sort(sort.StringSlice(in))
	fmt.Println(order)

	for _, m := range order {
		fmt.Printf("%q ", in[m])
	}
	fmt.Println()

	// Output:
	// [1 3 2 0]
	// "a" "b" "c" "z"
}

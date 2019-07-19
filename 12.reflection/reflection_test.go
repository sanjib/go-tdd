package reflection

import (
	"reflect"
	"testing"
)

func TestWalkTableDriven(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"struct with one string field", struct{ Name string }{"Sanjib"}, []string{"Sanjib"}},
		{"struct with two string fields",
			struct {
				Name string
				City string
			}{"Sanjib", "Chittagong"},
			[]string{"Sanjib", "Chittagong"}},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

func TestWalk(t *testing.T) {
	expected := "Sanjib"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d, want %d", len(got), 1)
	}
	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}

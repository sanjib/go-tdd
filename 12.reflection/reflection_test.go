package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

type CountryCode struct {
	Number  int
	Country string
}

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
			[]string{"Sanjib", "Chittagong"},
		},
		{"struct with two string fields",
			struct {
				Name string
				Age  int
			}{"Sanjib", 44},
			[]string{"Sanjib"},
		},
		{"nested fields",
			Person{"Sanjib", Profile{44, "Chittagong"}},
			[]string{"Sanjib", "Chittagong"},
		},
		{"pointer",
			&Person{"Sanjib", Profile{44, "Chittagong"}},
			[]string{"Sanjib", "Chittagong"},
		},
		{"slices",
			[]CountryCode{
				{45, "Denmark"},
				{49, "Germany"},
			},
			[]string{"Denmark", "Germany"},
		},
		{
			"arrays",
			[2]CountryCode{
				{45, "Denmark"},
				{49, "Germany"},
			},
			[]string{"Denmark", "Germany"},
		},
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

// TestWalkMap is in it's own function as map order is indeterministic
func TestWalkMap(t *testing.T) {
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"apple": "red",
			"mango": "green",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "green")
		assertContains(t, got, "red")
	})
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

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didnt", haystack, needle)
	}
}

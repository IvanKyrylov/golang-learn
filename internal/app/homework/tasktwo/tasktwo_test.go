package tasktwo

import "testing"

type testcase struct {
	name string
	s    string
	want string
}

func TestUnpackingString(t *testing.T) {
	tests := [...]testcase{
		{
			name: "one",
			s:    "a4bc2d5e",
			want: "aaaabccddddde",
		},
		{
			name: "two",
			s:    "abcd",
			want: "abcd",
		},
		{
			name: "three",
			s:    "45",
			want: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := UnpackingString(tc.s)
			if got != tc.want {
				t.Errorf("UnpackingString(%s) returns %s, want: %s", tc.s, got, tc.want)
			}
		})
	}
}

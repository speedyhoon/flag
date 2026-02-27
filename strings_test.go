package flag

import (
	"fmt"
	"testing"
)

func TestStrList_Set(t *testing.T) {
	tests := []struct {
		s    string
		want stringsValue
	}{
		{s: "", want: stringsValue{}},
		{s: " \n\t \t ", want: stringsValue{}},
		{s: " \n\t\v\a\b \t\v ", want: stringsValue{"\a\b"}},
		{s: "empty", want: stringsValue{"empty"}},
		{s: " empty ", want: stringsValue{"empty"}},
		{s: "\n\r\n\r\n\r\n\r\nempty\r \n\r\n\r", want: stringsValue{"empty"}},
		{s: " \n\t \t  empty\t\t ", want: stringsValue{"empty"}},
		{s: " \n\t\v\a\b \t\v  empty\v\t\a\b\t ", want: stringsValue{"\a\b \t\v  empty\v\t\a\b"}},
		{s: "empty,none", want: stringsValue{"empty", "none"}},
		{s: " \n\r \t  \fempty\t\f\t ,\f \n\t \r  \fnone\r\t \f", want: stringsValue{"empty", "none"}},
		{s: " \n\t\v\a \t\v  empty\v\t\a\t , \n\t\v\a \t\v  none\v\t\a\t ", want: stringsValue{"\a \t\v  empty\v\t\a", "\a \t\v  none\v\t\a"}},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", n), func(t *testing.T) {
			var got stringsValue
			if err := got.Set(tt.s); err != nil {
				t.Errorf("Set() error = %v, expected nil", err)
			}

			if len(got) != len(tt.want) {
				t.Errorf("Set() length = %d, want %d", len(got), len(tt.want))
			}

			for i, w := range tt.want {
				if got[i] != w {
					t.Errorf("Set()[%d]: `%s`, expected: `%s`", i, got[i], w)
				}
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		l    stringsValue
		want string
	}{
		{l: nil, want: ""},
		{l: stringsValue{}, want: ""},
		{l: stringsValue{""}, want: ""},
		{l: stringsValue{"empty"}, want: "empty"},
		{l: stringsValue{"\"empty, none\""}, want: `"empty, none"`},
		{l: stringsValue{"empty", "none"}, want: "empty, none"},
		{l: stringsValue{`"empty"`, `"none"`}, want: `"empty", "none"`},
		{l: stringsValue{"\n\t\v\a\b \t\v empty\n\r \t\v\t\a\b\t\n"}, want: "\n\t\v\a\b \t\v empty\n\r \t\v\t\a\b\t\n"},
		{l: stringsValue{
			"\n\t\v\a\b \t\v empty\n\r \t\v\t\a\b\t\n",
			"\n\t\v\a\b \t\v none\n\r \t\v\t\a\b\t\n",
		}, want: "\n\t\v\a\b \t\v empty\n\r \t\v\t\a\b\t\n, \n\t\v\a\b \t\v none\n\r \t\v\t\a\b\t\n"},
	}
	for n, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", n), func(t *testing.T) {
			got := tt.l.String()
			if got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

package flag

import (
	"strings"
)

// Sep is the default string separator used in stringsValue.Set.
var Sep = ","

// StringsVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a []string variable in which to store the value of the flag.
func (f *FlagSet) StringsVar(p *[]string, name string, value []string, usage string) {
	f.Var(newStringsValue(value, p), name, usage)
}

// StringsVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringsVar(p *[]string, name string, value []string, usage string) {
	CommandLine.Var(newStringsValue(value, p), name, usage)
}

// Strings defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSet) Strings(name string, value []string, usage string) (p []string) {
	f.StringsVar(&p, name, value, usage)
	return
}

// Strings defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func Strings(name string, value []string, usage string) []string {
	return CommandLine.Strings(name, value, usage)
}

// stringsValue implements the flag.Value interface.
type stringsValue []string

func newStringsValue(val []string, p *[]string) *stringsValue {
	*p = val
	return (*stringsValue)(p)
}

// Set slices s into all substrings separated by Sep and returns a slice of
// the substrings between those separators with leading and trailing whitespace removed.
func (s *stringsValue) Set(val string) error {
	var t stringsValue
	for _, item := range strings.Split(val, Sep) {
		item = strings.TrimSpace(item)
		if item != "" {
			t = append(t, item)
		}
	}
	*s = t // Clears any default value instead of appending to it.
	return nil
}

func (s *stringsValue) Get() any { return []string(*s) }

// String returns the string representation of the stringsValue, joining its elements with commas and a space.
func (s *stringsValue) String() string { return strings.Join(*s, ", ") }

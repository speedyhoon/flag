package flag

// StringVarOptional defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *FlagSet) StringVarOptional(p *string, name string, value string, usage string) *bool {
	newVal := newStringValueIsSpecified(value, p)
	f.Var(newVal, name, usage)
	return &newVal.IsSpecified
}

// StringVarOptional defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
// Returns true when the flag is listed even without a value.
func StringVarOptional(p *string, name string, value string, usage string) *bool {
	return CommandLine.StringVarOptional(p, name, value, usage)
}

// StringOptional defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSet) StringOptional(name string, value string, usage string) (*string, *bool) {
	p, bb := new(string), new(bool)
	f.StringVarOptional(p, name, value, usage)
	return p, bb
}

// StringOptional defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func StringOptional(name string, value string, usage string) (*string, *bool) {
	return CommandLine.StringOptional(name, value, usage)
}

type stringValueOptional struct {
	Value       *string
	IsSpecified bool
}

func newStringValueIsSpecified(val string, p *string) *stringValueOptional {
	*p = val
	return &stringValueOptional{Value: p}
}
func (s *stringValueOptional) Set(val string) error {
	*s.Value = val
	s.IsSpecified = true
	return nil
}

func (s *stringValueOptional) Get() any { return s.String() }

// IsBoolFlag allows flag without requiring a string value.
func (s *stringValueOptional) IsBoolFlag() bool { return true }

func (s *stringValueOptional) String() string {
	if s != nil && s.Value != nil {
		return *s.Value
	}
	return ""
}

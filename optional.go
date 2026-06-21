package flag

// StringVarOptional defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *FlagSet) StringVarOptional(p *string, isSet *bool, name, value, usage string) {
	newVal := newStringValueIsSpecified(value, p, isSet)
	f.Var(newVal, name, usage)
}

// StringVarOptional defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
// Returns true when the flag is listed even without a value.
func StringVarOptional(p *string, isSet *bool, name, value, usage string) {
	CommandLine.StringVarOptional(p, isSet, name, value, usage)
}

// StringOptional defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func (f *FlagSet) StringOptional(name, value, usage string) (*string, *bool) {
	p, t := new(string), new(bool)
	f.StringVarOptional(p, t, name, value, usage)
	return p, t
}

// StringOptional defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func StringOptional(name, value, usage string) (*string, *bool) {
	return CommandLine.StringOptional(name, value, usage)
}

type stringValueOptional struct {
	Value       *string
	IsSpecified *bool
}

func newStringValueIsSpecified(val string, p *string, isSet *bool) *stringValueOptional {
	*p = val
	return &stringValueOptional{Value: p, IsSpecified: isSet}
}

func (s *stringValueOptional) Set(val string) error {
	*s.IsSpecified = true
	if val != "true" {
		*s.Value = val
	}
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

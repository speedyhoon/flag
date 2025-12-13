package flag

import "strconv"

type uint16Value uint16

func newUint16Value(val uint16, p *uint16) *uint16Value {
	*p = val
	return (*uint16Value)(p)
}

func (i *uint16Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*i = uint16Value(v)
	return err
}

func (i *uint16Value) Get() any { return uint16(*i) }

func (i *uint16Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

// Uint16Var defines a uint16 flag with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the flag.
func (f *FlagSet) Uint16Var(p *uint16, name string, value uint16, usage string) {
	f.Var(newUint16Value(value, p), name, usage)
}

// Uint16Var defines a uint16 flag with specified name, default value, and usage string.
// The argument p points to a uint16 variable in which to store the value of the flag.
func Uint16Var(p *uint16, name string, value uint16, usage string) {
	CommandLine.Var(newUint16Value(value, p), name, usage)
}

// Uint16 defines a uint16 flag with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the flag.
func (f *FlagSet) Uint16(name string, value uint16, usage string) *uint16 {
	p := new(uint16)
	f.Uint16Var(p, name, value, usage)
	return p
}

// Uint16 defines a uint16 flag with specified name, default value, and usage string.
// The return value is the address of a uint16 variable that stores the value of the flag.
func Uint16(name string, value uint16, usage string) *uint16 {
	return CommandLine.Uint16(name, value, usage)
}

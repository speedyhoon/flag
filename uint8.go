package flag

import "strconv"

type uint8Value uint8

func newUint8Value(val uint8, p *uint8) *uint8Value {
	*p = val
	return (*uint8Value)(p)
}

func (i *uint8Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*i = uint8Value(v)
	return err
}

func (i *uint8Value) Get() any { return uint8(*i) }

func (i *uint8Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

// Uint8Var defines a uint8 flag with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the flag.
func (f *FlagSet) Uint8Var(p *uint8, name string, value uint8, usage string) {
	f.Var(newUint8Value(value, p), name, usage)
}

// Uint8Var defines a uint8 flag with specified name, default value, and usage string.
// The argument p points to a uint8 variable in which to store the value of the flag.
func Uint8Var(p *uint8, name string, value uint8, usage string) {
	CommandLine.Var(newUint8Value(value, p), name, usage)
}

// Uint8 defines a uint8 flag with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the flag.
func (f *FlagSet) Uint8(name string, value uint8, usage string) *uint8 {
	p := new(uint8)
	f.Uint8Var(p, name, value, usage)
	return p
}

// Uint8 defines a uint8 flag with specified name, default value, and usage string.
// The return value is the address of a uint8 variable that stores the value of the flag.
func Uint8(name string, value uint8, usage string) *uint8 {
	return CommandLine.Uint8(name, value, usage)
}

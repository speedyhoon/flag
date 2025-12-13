package flag

import "strconv"

type uint32Value uint32

func newUint32Value(val uint32, p *uint32) *uint32Value {
	*p = val
	return (*uint32Value)(p)
}

func (i *uint32Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*i = uint32Value(v)
	return err
}

func (i *uint32Value) Get() any { return uint32(*i) }

func (i *uint32Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

// Uint32Var defines a uint32 flag with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the flag.
func (f *FlagSet) Uint32Var(p *uint32, name string, value uint32, usage string) {
	f.Var(newUint32Value(value, p), name, usage)
}

// Uint32Var defines a uint32 flag with specified name, default value, and usage string.
// The argument p points to a uint32 variable in which to store the value of the flag.
func Uint32Var(p *uint32, name string, value uint32, usage string) {
	CommandLine.Var(newUint32Value(value, p), name, usage)
}

// Uint32 defines a uint32 flag with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the flag.
func (f *FlagSet) Uint32(name string, value uint32, usage string) *uint32 {
	p := new(uint32)
	f.Uint32Var(p, name, value, usage)
	return p
}

// Uint32 defines a uint32 flag with specified name, default value, and usage string.
// The return value is the address of a uint32 variable that stores the value of the flag.
func Uint32(name string, value uint32, usage string) *uint32 {
	return CommandLine.Uint32(name, value, usage)
}

package flag

import "strconv"

type int8Value int8

func newInt8Value(val int8, p *int8) *int8Value {
	*p = val
	return (*int8Value)(p)
}

func (i *int8Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*i = int8Value(v)
	return err
}

func (i *int8Value) Get() any { return int8(*i) }

func (i *int8Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

// Int8Var defines an int8 flag with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the flag.
func (f *FlagSet) Int8Var(p *int8, name string, value int8, usage string) {
	f.Var(newInt8Value(value, p), name, usage)
}

// Int8Var defines an int8 flag with specified name, default value, and usage string.
// The argument p points to an int8 variable in which to store the value of the flag.
func Int8Var(p *int8, name string, value int8, usage string) {
	CommandLine.Var(newInt8Value(value, p), name, usage)
}

// Int8 defines an int8 flag with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the flag.
func (f *FlagSet) Int8(name string, value int8, usage string) *int8 {
	p := new(int8)
	f.Int8Var(p, name, value, usage)
	return p
}

// Int8 defines an int8 flag with specified name, default value, and usage string.
// The return value is the address of an int8 variable that stores the value of the flag.
func Int8(name string, value int8, usage string) *int8 {
	return CommandLine.Int8(name, value, usage)
}

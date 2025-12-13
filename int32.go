package flag

import "strconv"

type int32Value int32

func newInt32Value(val int32, p *int32) *int32Value {
	*p = val
	return (*int32Value)(p)
}

func (i *int32Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	if err != nil {
		err = numError(err)
	}
	*i = int32Value(v)
	return err
}

func (i *int32Value) Get() any { return int32(*i) }

func (i *int32Value) String() string { return strconv.FormatUint(uint64(*i), 10) }

// Int32Var defines an int32 flag with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the flag.
func (f *FlagSet) Int32Var(p *int32, name string, value int32, usage string) {
	f.Var(newInt32Value(value, p), name, usage)
}

// Int32Var defines an int32 flag with specified name, default value, and usage string.
// The argument p points to an int32 variable in which to store the value of the flag.
func Int32Var(p *int32, name string, value int32, usage string) {
	CommandLine.Var(newInt32Value(value, p), name, usage)
}

// Int32 defines an int32 flag with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the flag.
func (f *FlagSet) Int32(name string, value int32, usage string) *int32 {
	p := new(int32)
	f.Int32Var(p, name, value, usage)
	return p
}

// Int32 defines an int32 flag with specified name, default value, and usage string.
// The return value is the address of an int32 variable that stores the value of the flag.
func Int32(name string, value int32, usage string) *int32 {
	return CommandLine.Int32(name, value, usage)
}

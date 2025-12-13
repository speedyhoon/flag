package flag

import (
	"errors"
	"net"
)

type ipValue net.IP

func newIPValue(val net.IP, p *net.IP) *ipValue {
	*p = val
	return (*ipValue)(p)
}

func (i *ipValue) Set(s string) error {
	v := net.ParseIP(s)
	if i == nil {
		return errors.New("could not parse IP")
	}
	*i = ipValue(v)
	return nil
}

func (i *ipValue) Get() any { return net.IP(*i) }

func (i *ipValue) String() string {
	if i == nil {
		return ""
	}
	return net.IP(*i).String()
}

// IPVar defines an IP flag with specified name, default value, and usage string.
// The argument p points to an IP variable in which to store the value of the flag.
func (f *FlagSet) IPVar(p *net.IP, name string, value net.IP, usage string) {
	f.Var(newIPValue(value, p), name, usage)
}

// IPVar defines an IP flag with specified name, default value, and usage string.
// The argument p points to an IP variable in which to store the value of the flag.
func IPVar(p *net.IP, name string, value net.IP, usage string) {
	CommandLine.Var(newIPValue(value, p), name, usage)
}

// IP defines an IP flag with specified name, default value, and usage string.
// The return value is the address of an IP variable that stores the value of the flag.
func (f *FlagSet) IP(name string, value net.IP, usage string) *net.IP {
	p := new(net.IP)
	f.IPVar(p, name, value, usage)
	return p
}

// IP defines an IP flag with specified name, default value, and usage string.
// The return value is the address of an IP variable that stores the value of the flag.
func IP(name string, value net.IP, usage string) *net.IP {
	return CommandLine.IP(name, value, usage)
}

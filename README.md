# flag 🏁

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/flag.svg)](https://pkg.go.dev/github.com/speedyhoon/flag)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/flag)](https://goreportcard.com/report/github.com/speedyhoon/flag)

An expansion upon Go's built-in package [`flag`](https://pkg.go.dev/flag) to support more types. It is a fully compatible, direct replacement.

## Example:

```go
package main

import (
	"fmt"
	"github.com/speedyhoon/flag"
	"net"
)

func main() {
	var o Options
	flag.Float32Var(&o.f32, "x", -180.0087, "coordinate")
	flag.Int8Var(&o.i8, "q", 5, "quantity")
	flag.Int16Var(&o.i16, "v", -32567, "inventory")
	flag.Int32Var(&o.i32, "r", '😀', "rune")
	flag.IPVar(&o.ip, "ip", net.IPv4(127, 0, 0, 1), "ip address")
	flag.Uint8Var(&o.u8, "o", 'Y', "enum option")
	flag.Uint16Var(&o.u16, "p", 80, "network port")
	flag.Uint32Var(&o.u32, "z", 16777216, "maximum size allowed")
	flag.StringsVar(&o.s, "a", []string{"panda", "tiger", "monkey", "viper"}, "list of animals")
	flag.StringVarOptional(&o.err, &o.isE, "set", "io.EOF", "default error")
	flag.Parse()

	fmt.Println(o)
}

type Options struct {
	f32 float32
	i8  int8
	i16 int16
	i32 int32
	ip  net.IP
	u8  uint8
	u16 uint16
	u32 uint32
	s   []string
	err string
	isE bool
}

```

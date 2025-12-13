# flag

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/flag.svg)](https://pkg.go.dev/github.com/speedyhoon/flag)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/flag)](https://goreportcard.com/report/github.com/speedyhoon/flag)

An expansion upon Go's built-in package `flag` to support more types.

## Example:

```go
package main

import (
	"fmt"
	"github.com/speedyhoon/flag"
	"net"
)

func main() {
	var t Types
	flag.Float32Var(&t.f32, "x", -180.0087, "coordinate")
	flag.Int8Var(&t.i8, "q", 5, "quantity")
	flag.Int16Var(&t.i16, "v", -32567, "inventory")
	flag.Int32Var(&t.i32, "r", '😀', "rune")
	flag.IPVar(&t.ip, "ip", net.IPv4(127, 0, 0, 1), "ip address")
	flag.Uint8Var(&t.u8, "o", 'Y', "enum option")
	flag.Uint16Var(&t.u16, "p", 80, "network port")
	flag.Uint32Var(&t.u32, "z", 16777216, "maximum size allowed")
	flag.Parse()

	fmt.Println(t)
}

type Types struct {
	f32  float32
	i8   int8
	i16  int16
	i32  int32
	ip   net.IP
	u8   uint8
	u16  uint16
	u32  uint32
}
```

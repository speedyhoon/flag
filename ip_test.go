package flag_test

import (
	"bytes"
	"github.com/speedyhoon/flag"
	"net"
	"testing"
)

func testParseIP(f *flag.FlagSet, t *testing.T) {
	if f.Parsed() {
		t.Error("f.Parse() = true before Parse")
	}
	ipFlag := f.IP("ip", net.IPv4(127, 0, 0, 1), "IP address")
	extra := "one-extra-argument"
	args := []string{
		"-ip=10.1.2.4",
		extra,
	}
	if err := f.Parse(args); err != nil {
		t.Fatal(err)
	}
	if !f.Parsed() {
		t.Error("f.Parse() = false after Parse")
	}
	if !bytes.Equal(*ipFlag, net.IPv4(10, 1, 2, 4)) {
		t.Error("ip flag should be 10.1.2.4, is ", *ipFlag)
	}
	if len(f.Args()) != 1 {
		t.Error("expected one argument, got", len(f.Args()))
	} else if f.Args()[0] != extra {
		t.Errorf("expected argument %q got %q", extra, f.Args()[0])
	}
}

func TestParseIP(t *testing.T) {
	flag.ResetForTesting(func() { t.Error("bad parse") })
	testParseIP(flag.CommandLine, t)
}

func TestFlagSetParseIP(t *testing.T) {
	testParseIP(flag.NewFlagSet("IPv4", flag.ContinueOnError), t)
}

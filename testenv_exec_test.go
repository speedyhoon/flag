// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flag_test

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"testing"
)

// MustHaveExec checks that the current system can start new processes
// using os.StartProcess or (more commonly) exec.Command.
// If not, MustHaveExec calls t.Skip with an explanation.
//
// On some platforms MustHaveExec checks for exec support by re-executing the
// current executable, which must be a binary built by 'go test'.
// We intentionally do not provide a HasExec function because of the risk of
// inappropriate recursion in TestMain functions.
//
// To check for exec support outside of a test, just try to exec the command.
// If exec is not supported, testenv.SyscallIsNotSupported will return true
// for the resulting error.
func MustHaveExec(t testing.TB) {
	if err := tryExec(); err != nil {
		msg := fmt.Sprintf("cannot exec subprocess on %s/%s: %v", runtime.GOOS, runtime.GOARCH, err)
		if t == nil {
			panic(msg)
		}
		t.Helper()
		t.Skip("skipping test:", msg)
	}
}

var tryExec = sync.OnceValue(func() error {
	switch runtime.GOOS {
	case "wasip1", "js", "ios":
	default:
		// Assume that exec always works on non-mobile platforms and Android.
		return nil
	}

	// ios has an exec syscall but on real iOS devices it might return a
	// permission error. In an emulated environment (such as a Corellium host)
	// it might succeed, so if we need to exec we'll just have to try it and
	// find out.
	//
	// As of 2023-04-19 wasip1 and js don't have exec syscalls at all, but we
	// may as well use the same path so that this branch can be tested without
	// an ios environment.

	if !testing.Testing() {
		// This isn't a standard 'go test' binary, so we don't know how to
		// self-exec in a way that should succeed without side effects.
		// Just forget it.
		return errors.New("can't probe for exec support with a non-test executable")
	}

	// We know that this is a test executable. We should be able to run it with a
	// no-op flag to check for overall exec support.
	exe, err := exePath()
	if err != nil {
		return fmt.Errorf("can't probe for exec support: %w", err)
	}
	cmd := exec.Command(exe, "-test.list=^$")
	cmd.Env = origEnv
	return cmd.Run()
})

var exePath = sync.OnceValues(func() (string, error) {
	return os.Executable()
})

// +build linux,!android

package signals

import "syscall"

// Missing signals in Golang
const (
	SIGINFO = syscall.SIGPWR
)

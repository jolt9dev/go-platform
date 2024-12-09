//go:build aix || dragonfly || freebsd || hurd || illumos || linux || netbsd || openbsd || plan9 || solaris || zos

package platform

func IsDarwin() bool {
	return false
}

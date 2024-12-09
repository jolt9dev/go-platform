package platform_test

import (
	"runtime"
	"testing"

	"github.com/jolt9dev/go-platform"
	"github.com/stretchr/testify/assert"
)

func TestOs(t *testing.T) {
	assert.Equal(t, platform.OS, runtime.GOOS)
}

func TestWindows(t *testing.T) {
	if runtime.GOOS == "windows" {
		assert.True(t, platform.IsWindows())
		return
	}

	assert.False(t, platform.IsWindows())
}

func TestDarwin(t *testing.T) {
	if runtime.GOOS == "darwin" {
		assert.True(t, platform.IsDarwin())
		return
	}

	assert.False(t, platform.IsDarwin())
}

func TestLinux(t *testing.T) {
	if runtime.GOOS == "linux" {
		assert.True(t, platform.IsLinux())
		return
	}

	assert.False(t, platform.IsLinux())
}

func TestPosix(t *testing.T) {
	if runtime.GOOS == "windows" {
		assert.False(t, platform.IsPosix())
		return
	}

	assert.True(t, platform.IsPosix())
}

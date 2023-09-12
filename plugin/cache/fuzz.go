//go:build gofuzz

package cache

import (
	"github.com/khulnasoft-lab/dnserver/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	return fuzz.Do(New(), data)
}

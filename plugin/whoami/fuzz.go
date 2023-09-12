//go:build gofuzz

package whoami

import (
	"github.com/khulnasoft-lab/dnserver/plugin/pkg/fuzz"
)

// Fuzz fuzzes cache.
func Fuzz(data []byte) int {
	w := Whoami{}
	return fuzz.Do(w, data)
}

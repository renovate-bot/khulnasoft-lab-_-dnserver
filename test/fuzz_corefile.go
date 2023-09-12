//go:build gofuzz

package test

// Fuzz fuzzes a corefile.
func Fuzz(data []byte) int {
	_, _, _, err := DNServerServerAndPorts(string(data))
	if err != nil {
		return 1
	}
	return 0
}

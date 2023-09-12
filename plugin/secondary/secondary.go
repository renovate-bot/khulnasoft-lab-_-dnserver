// Package secondary implements a secondary plugin.
package secondary

import "github.com/khulnasoft-lab/dnserver/plugin/file"

// Secondary implements a secondary plugin that allows DNServer to retrieve (via AXFR)
// zone information from a primary server.
type Secondary struct {
	file.File
}

// Name implements the Handler interface.
func (s Secondary) Name() string { return "secondary" }

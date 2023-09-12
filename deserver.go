package main

import (
	_ "github.com/khulnasoft-lab/dnserver/core/plugin" // Plug in DNServer.
	"github.com/khulnasoft-lab/dnserver/core"
)

func main() {
	core.Run()
}

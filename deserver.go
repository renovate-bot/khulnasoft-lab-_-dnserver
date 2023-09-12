package main

import (
	_ "github.com/khulnasoft-lab/dnserver/core/plugin" // Plug in DNServer.
	"github.com/khulnasoft-lab/dnserver/coremain"
)

func main() {
	coremain.Run()
}

package view

import (
	"context"

	"github.com/khulnasoft-lab/dnserver/plugin/metadata"
	"github.com/khulnasoft-lab/dnserver/request"
)

// Metadata implements the metadata.Provider interface.
func (v *View) Metadata(ctx context.Context, state request.Request) context.Context {
	metadata.SetValueFunc(ctx, "view/name", func() string {
		return v.viewName
	})
	return ctx
}

package main

import (
	"github.com/hashicorp/terraform/builtin/providers/opc"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: opc.Provider,
	})
}

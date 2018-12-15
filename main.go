package main

import (
	"github.com/doubret/terraform-provider-xiiot/xiiot"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: xiiot.Provider,
	}

	plugin.Serve(&opts)
}

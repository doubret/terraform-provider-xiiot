package main

import (
	"github.com/doubret/terraform-provider-xiiot/xiiot"
	"github.com/hashicorp/terraform/plugin"
	"github.com/doubret/terraform-provider-xiiot/xiiot/client/client"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: xiiot.Provider,
	}

	resp, err := apiclient.Default.Operations.All(operations.AllParams{})
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%#v\n", resp.Payload)

	plugin.Serve(&opts)
}
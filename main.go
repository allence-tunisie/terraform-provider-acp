package main

import (
	"github.com/allence-tunisie/terraform-provider-acp/acp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: acp.Provider,
	})
}

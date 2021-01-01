package acp

import (
	"context"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
		},
		ResourcesMap: map[string]*schema.Resource{
		},
		DataSourcesMap: map[string]*schema.Resource{
		},
		ConfigureContextFunc: providerConfigure,

	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	var client any.Any
	return client,diags
}


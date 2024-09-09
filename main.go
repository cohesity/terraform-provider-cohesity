package main

//go:generate terraform fmt -recursive ./examples/

//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name cohesity

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/terraform-providers/terraform-provider-cohesity/cohesity"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cohesity.Provider})
}

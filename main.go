package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/alexey-mylnikov/terraform-provider-runtime/internal/provider"
)

// Generate the Terraform provider documentation using `tfplugindocs`:
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address:         "registry.terraform.io/alexey-mylnikov/runtime",
		Debug:           debug,
		ProtocolVersion: 5,
	})

	if err != nil {
		log.Fatal(err)
	}
}

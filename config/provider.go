/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/chelala/provider-azureextraupjet/config/azurerm_eventgrid_system_topic_event_subscription"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "azureextraupjet"
	modulePath     = "github.com/chelala/provider-azureextraupjet"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("idp.microserfin.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		azurerm_eventgrid_system_topic_event_subscription.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

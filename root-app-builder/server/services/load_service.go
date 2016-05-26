package services

import (
	"log"

	"github.com/eogile/agilestack-utils/plugins/components"
	"github.com/eogile/agilestack-utils/plugins/registration"
)

func LoadApplication() (configurations []registration.PluginConfiguration, appComponents *components.Components, err error) {
	configurations, err = registration.ListRoutesAndReducers()
	if err != nil {
		log.Println("Error while loading configurations from Consul", err)
		return
	}

	appComponents, err = components.GetComponents()
	if err != nil {
		log.Println("Error while loading app components from Consul", err)
		return
	}

	return
}

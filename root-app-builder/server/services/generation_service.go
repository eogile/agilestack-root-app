package services

import (
	"log"

	appFiles "github.com/eogile/agilestack-root-app/root-app-builder/server/files"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/generators"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/utils/npm"
	"github.com/eogile/agilestack-utils/plugins/components"
	"github.com/eogile/agilestack-utils/plugins/registration"
)

func generateFiles(configurations []registration.PluginConfiguration, components *components.Components) error {
	err := generateReducersFile(configurations)
	if err != nil {
		log.Printf("Error while generating the reducers file: %v\n", err)
		return err
	}
	err = generateRoutesFile(configurations)
	if err != nil {
		log.Printf("Error while generating the routes file: %v\n", err)
	}

	err = generateComponentsFile(components)
	if err != nil {
		log.Printf("Error while generating the components file: %v\n", err)
	}
	return err
}

func generateReducersFile(configurations []registration.PluginConfiguration) error {
	return generators.GenerateReducersFile(
		configurations,
		appFiles.ReducersFile)
}

func generateRoutesFile(configurations []registration.PluginConfiguration) error {
	return generators.GenerateRoutesFile(
		configurations,
		appFiles.RoutesFile)
}

func generateComponentsFile(components *components.Components) error {
	return generators.GenerateComponentsFile(components, appFiles.ComponentsFile)
}

func GenerateApplication(configurations []registration.PluginConfiguration, components *components.Components) error {

	/*
	 * Generating the routes and reducers file.
	 */
	err := generateFiles(configurations, components)
	if err != nil {
		return err
	}

	/*
	 * Linking the NPM package.
	 */
	for _, config := range configurations {
		err = npm.LinkPackage(config.PluginName)
		if err != nil {
			log.Println("Error while NPM package linking:", err)
			return err
		}
	}

	return nil
}

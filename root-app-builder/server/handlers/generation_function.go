package handlers

import (
	"log"
	"net/http"

	appFiles "github.com/eogile/agilestack-root-app/root-app-builder/server/files"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/generators"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/services"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/utils/npm"
	"github.com/eogile/agilestack-utils/plugins/registration"
)

func generateFiles(configurations []registration.PluginConfiguration) error {
	err := generateReducersFile(configurations)
	if err != nil {
		log.Printf("Error while generating the reducers file: %v\n", err)
		return err
	}
	err = generateRoutesFile(configurations)
	if err != nil {
		log.Printf("Error while generating the routes file: %v\n", err)
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

func BuildApplication(w http.ResponseWriter, configurations []registration.PluginConfiguration) {
	/*
	 * Generating the routes and reducers file.
	 */
	err := generateFiles(configurations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/*
	 * Linking the NPM package.
	 */
	for _, config := range configurations {
		err = npm.LinkPackage(config.PluginName)
		if err != nil {
			log.Println("Error while NPM package linking:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	/*
	 * Compiling the JavaScript application.
	 */
	err = services.BuildApplication()
	if err != nil {
		log.Println("Error while building compilation:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

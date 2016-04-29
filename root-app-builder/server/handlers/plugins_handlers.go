package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	appFiles "github.com/eogile/agilestack-root-app/root-app-builder/server/files"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/generators"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/models"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/repository"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/services"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/utils/npm"
)

var pluginConfigurationsRepo = repository.GetPluginConfigurationsRepository()

func HandlePluginsEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Treating request \"%s %s\"", r.Method, r.URL)
	switch r.Method {
	case "POST":
		handlePluginsPost(w, r)
	default:
		log.Printf("Request %s - Method not allowed : %s", r.URL, r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func handlePluginsPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	var config models.PluginConfiguration
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error while parsing request body", err)
		return
	}
	pluginConfigurationsRepo.Add(config)

	/*
	 * Generating the routes and reducers file.
	 */
	err = generateFiles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/*
	 * Linking the NPM package.
	 */
	err = npm.LinkPackage(config.ModuleName)
	if err != nil {
		log.Println("Error while NPM package linking:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
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

func generateFiles() error {
	err := generateReducersFile()
	if err != nil {
		log.Printf("Error while generating the reducers file: %v\n", err)
		return err
	}
	err = generateRoutesFile()
	if err != nil {
		log.Printf("Error while generating the routes file: %v\n", err)
	}
	return err
}

func generateReducersFile() error {
	return generators.GenerateReducersFile(
		pluginConfigurationsRepo.GetReducers(),
		appFiles.ReducersFile)
}

func generateRoutesFile() error {
	return generators.GenerateRoutesFile(
		pluginConfigurationsRepo.GetRoutes(),
		appFiles.RoutesFile)
}

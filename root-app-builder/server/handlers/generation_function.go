package handlers

import (
	"log"
	"net/http"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/services"
	"github.com/eogile/agilestack-utils/plugins/components"
	"github.com/eogile/agilestack-utils/plugins/registration"
)

func BuildApplicationHandler(w http.ResponseWriter,
	configurations []registration.PluginConfiguration,
	components *components.Components) {

	err := services.GenerateApplication(configurations, components)
	if err != nil {
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

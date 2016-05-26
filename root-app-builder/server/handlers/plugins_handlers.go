package handlers

import (
	"log"
	"net/http"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/services"
	"github.com/eogile/agilestack-utils/plugins/components"
	"github.com/eogile/agilestack-utils/plugins/registration"
)

type registrationHandler struct {
	delegateFunction func(http.ResponseWriter, []registration.PluginConfiguration, *components.Components)
}

func NewGenerationHandler(generationFunction func(http.ResponseWriter, []registration.PluginConfiguration, *components.Components)) *registrationHandler {
	return &registrationHandler{
		delegateFunction: generationFunction,
	}
}

func (h *registrationHandler) HandlePluginsEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Treating request \"%s %s\"", r.Method, r.URL)
	switch r.Method {
	case "POST":
		h.handlePluginsPost(w, r)
	default:
		log.Printf("Request %s - Method not allowed : %s", r.URL, r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *registrationHandler) handlePluginsPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	configurations, appComponents, err := services.LoadApplication()
	if err != nil {
		return
	}

	/*
	 * Calling the delegate function
	 */
	h.delegateFunction(w, configurations, appComponents)
}

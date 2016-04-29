package handlers

import (
	"log"
	"net/http"

	"encoding/json"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/repository"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/models"
)

/*
 * The repository managing the menu entries persistence.
 */
var repo = repository.GetRepository()

func HandleMenuEntriesEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Treating request \"%s %s\"", r.Method, r.URL)

	switch r.Method {
	case "GET":
		handleMenuGET(w, r)
	case "POST":
		handleMenuPOST(w, r)
	default:
		log.Printf("Request %s - Method not allowed : %s", r.URL, r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

/*
 * swagger:route GET / listUserEntries
 *
 * List all the menu entries.
 *
 * Produces:application/json
 *
 * Schemes: http, https
 *
 * Responses:
 *   200: listOfMenuEntries
 *   500: errorResponse
 */
func handleMenuGET(w http.ResponseWriter, r *http.Request) {
	entries, err := repo.FindAll()
	w.Header().Set("Content-Type", "application/json")

	/*
	 * Adding CORS headers for Swagger UI.
	 */
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		bytes, err := json.Marshal(entries)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}

func handleMenuPOST(w http.ResponseWriter, r *http.Request) {
	var entries []models.MenuEntry
	err := json.NewDecoder(r.Body).Decode(&entries)

	if err != nil {
		log.Println("Error while reading request body:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = repo.Save(entries)
	if err != nil {
		log.Println("Error while saving menu entries:", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)

	}
}

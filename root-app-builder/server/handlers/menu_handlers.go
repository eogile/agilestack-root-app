package handlers

import (
	"log"
	"net/http"
	"sort"

	"encoding/json"

	"github.com/eogile/agilestack-utils/plugins/menu"
	"strings"
)

func HandleMenuEntriesEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Treating request \"%s %s\"", r.Method, r.URL)

	switch r.Method {
	case "GET":
		handleMenuGET(w, r)
	default:
		log.Printf("Request %s - Method not allowed : %s", r.URL, r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

/*
 * swagger:route GET / listUserEntries
 *
 * List all the menu entries sorted by weight desc and then name asc.
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
	entries, err := menu.ListMenus()
	w.Header().Set("Content-Type", "application/json")

	/*
	 * Adding CORS headers for Swagger UI.
	 */
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		bytes, err := json.Marshal(sortMenuEntries(entries))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		}
	}
}

func sortMenuEntries(menus []menu.Menu) []menu.MenuEntry {
	entries := make([]menu.MenuEntry, 0)
	for _, menu := range menus {
		entries = append(entries, menu.Entries...)
	}

	sortableData := &SortableData{
		menuEntries: entries,
	}
	sort.Sort(sortableData)
	return sortableData.menuEntries
}

type SortableData struct {
	menuEntries []menu.MenuEntry
}

func (data *SortableData) Len() int {
	return len(data.menuEntries)
}

func (data *SortableData) Less(i, j int) bool {
	if data.menuEntries[i].Weight != data.menuEntries[j].Weight {
		return data.menuEntries[i].Weight < data.menuEntries[j].Weight
	}
	return strings.Compare(data.menuEntries[i].Name, data.menuEntries[j].Name) < 0
}

func (data *SortableData) Swap(i, j int) {
	valueI := data.menuEntries[i]
	data.menuEntries[i] = data.menuEntries[j]
	data.menuEntries[j] = valueI
}
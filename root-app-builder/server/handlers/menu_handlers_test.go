package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	testUtils "github.com/eogile/agilestack-root-app/root-app-builder/server/testing"
	"os"
	"github.com/eogile/agilestack-utils/plugins/menu"
	"encoding/json"
)

var (
	menu1 = menu.Menu{
		PluginName: "MyWonderfulPlugin",
		Entries: []menu.MenuEntry{
			menu.MenuEntry{
				Name:   "Entry 01",
				Route:  "/entry-1",
				Weight: 10,
				Entries: []menu.MenuEntry{},
			},
			menu.MenuEntry{
				Name:   "Entry 05",
				Route:  "/entry-2",
				Weight: 1,
				Entries: []menu.MenuEntry{
					menu.MenuEntry{
						Name:   "Entry 2.1",
						Route:  "/entry-2-1",
						Weight: 10,
						Entries: []menu.MenuEntry{},
					},
				},
			},
		},
	}

	menu2 = menu.Menu{
		PluginName: "plugin_2",
		Entries: []menu.MenuEntry{
			menu.MenuEntry{
				Name:   "Entry 04",
				Route:  "/entry-4",
				Weight: 1,
				Entries: []menu.MenuEntry{},
			},
		},
	}
)

func init() {
	os.Setenv("SHARED_MODULES_DIRECTORY", "../../web_modules/")
	os.Setenv("HTTP_DIRECTORY", "../../build/")
	os.Setenv("SOURCES_DIRECTORY", "../../")
}

func TestGetMenuEntries(t *testing.T) {
	testUtils.DeleteAllStoreEntries(t)

	err := menu.StoreMenu(&menu1)
	require.Nil(t, err)
	err = menu.StoreMenu(&menu2)
	require.Nil(t, err)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleMenuEntriesEndpoint)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	mux.ServeHTTP(writer, request)

	/*
	 * Checking the HTTP status.
	 */
	assert.Equal(t, 200, writer.Code, "Invalid HTTP status")

	/*
	 * Checking the content type.
	 */
	contentType := writer.Header().Get("Content-Type")
	assert.Equal(t, "application/json", contentType, "Invalid Content-Type header")

	var menuEntries []menu.MenuEntry
	err = json.Unmarshal(writer.Body.Bytes(), &menuEntries)
	require.Nil(t, err, "Error while reading response body")
	require.Equal(t, 3, len(menuEntries), "Invalid number of menu entries")

	/*
	 * Top level entries are sorted by "weight asc" and then "name asc".
	 */
	validateEntry(t, menu2.Entries[0], menuEntries[0])
	validateEntry(t, menu1.Entries[1], menuEntries[1])
	validateEntry(t, menu1.Entries[0], menuEntries[2])
}

func TestGetMenuEntries_NoMenu(t *testing.T) {
	testUtils.DeleteAllStoreEntries(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleMenuEntriesEndpoint)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	mux.ServeHTTP(writer, request)

	/*
	 * Checking the HTTP status.
	 */
	assert.Equal(t, 200, writer.Code, "Invalid HTTP status")

	/*
	 * Checking the content type.
	 */
	contentType := writer.Header().Get("Content-Type")
	assert.Equal(t, "application/json", contentType, "Invalid Content-Type header")

	var menuEntries []menu.MenuEntry
	err := json.Unmarshal(writer.Body.Bytes(), &menuEntries)
	require.Nil(t, err, "Error while reading response body")
	require.Equal(t, 0, len(menuEntries), "Invalid number of menu entries")
}

func validateEntry(t *testing.T, expectedEntry menu.MenuEntry, resultEntry menu.MenuEntry) {
	assert.Equal(t, expectedEntry.Name, resultEntry.Name)
	assert.Equal(t, expectedEntry.Route, resultEntry.Route)
	assert.Equal(t, expectedEntry.Weight, resultEntry.Weight)
	require.Equal(t, len(expectedEntry.Entries), len(resultEntry.Entries))

	for i, expectedSubEntry := range expectedEntry.Entries {
		resultSubEntry := resultEntry.Entries[i]
		validateEntry(t, expectedSubEntry, resultSubEntry)
	}
}
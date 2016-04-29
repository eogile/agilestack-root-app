package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/handlers"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"bytes"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/repository"
	testUtils "github.com/eogile/agilestack-root-app/root-app-builder/server/testing"
	"os"
)

func init() {
	os.Setenv("SHARED_MODULES_DIRECTORY", "../../web_modules/")
	os.Setenv("HTTP_DIRECTORY", "../../build/")
	os.Setenv("SOURCES_DIRECTORY", "../../")
}

func TestGetMenuEntries(t *testing.T) {

	testUtils.SetUp(t)

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

	/*
	 * Checking the response body.
	 */
	var menuEntries []models.MenuEntry
	err := json.Unmarshal(writer.Body.Bytes(), &menuEntries)
	require.Nil(t, err, "Error while reading response body")
	require.Equal(t, 3, len(menuEntries), "Invalid number of menu entries")

	/*
	 * First menu entry
	 */
	assert.Equal(t, "plugin1", menuEntries[0].PluginName)
	assert.Equal(t, "menu 1", menuEntries[0].Name)
	assert.Equal(t, "route1", menuEntries[0].Route)
	assert.Equal(t, 10, menuEntries[0].Weight)
	assert.NotNil(t, menuEntries[0].ID)

	/*
	 * Second menu entry
	 */
	assert.Equal(t, "plugin2", menuEntries[1].PluginName)
	assert.Equal(t, "menu 3", menuEntries[1].Name)
	assert.Equal(t, "route3", menuEntries[1].Route)
	assert.Equal(t, 10, menuEntries[1].Weight)
	assert.NotNil(t, menuEntries[1].ID)

	/*
	 * Third menu entry
	 */
	assert.Equal(t, "plugin1", menuEntries[2].PluginName)
	assert.Equal(t, "menu 2", menuEntries[2].Name)
	assert.Equal(t, "route2", menuEntries[2].Route)
	assert.Equal(t, 5, menuEntries[2].Weight)
	assert.NotNil(t, menuEntries[2].ID)
}

func TestPostMenuEntries(t *testing.T) {
	testUtils.SetUp(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleMenuEntriesEndpoint)

	writer := httptest.NewRecorder()

	entries := []models.MenuEntry{
		{
			PluginName: "plugin2",
			Name:       "Todo list application",
			Route:      "/todo-list",
			Weight:     10,
		},
		{
			PluginName: "plugin2",
			Name:       "1 - plugin2 Menu entry1",
			Route:      "/plugin2/feature1",
			Weight:     5,
		},
	}

	jsonBytes, err := json.Marshal(entries)
	require.Nil(t, err)
	request, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonBytes))

	mux.ServeHTTP(writer, request)

	assert.Equal(t, 201, writer.Code, "Invalid HTTP status")

	menuEntries, _ := repository.GetRepository().FindAll()
	assert.Equal(t, 4, len(menuEntries))

	/*
	 * First menu entry
	 */
	assert.Equal(t, "plugin1", menuEntries[0].PluginName)
	assert.Equal(t, "menu 1", menuEntries[0].Name)
	assert.Equal(t, "route1", menuEntries[0].Route)
	assert.Equal(t, 10, menuEntries[0].Weight)
	assert.NotNil(t, menuEntries[0].ID)

	/*
	 * Second menu entry
	 */
	assert.Equal(t, "plugin2", menuEntries[1].PluginName)
	assert.Equal(t, "Todo list application", menuEntries[1].Name)
	assert.Equal(t, "/todo-list", menuEntries[1].Route)
	assert.Equal(t, 10, menuEntries[1].Weight)
	assert.NotNil(t, menuEntries[1].ID)

	/*
	 * Third menu entry
	 */
	assert.Equal(t, "plugin2", menuEntries[2].PluginName)
	assert.Equal(t, "1 - plugin2 Menu entry1", menuEntries[2].Name)
	assert.Equal(t, "/plugin2/feature1", menuEntries[2].Route)
	assert.Equal(t, 5, menuEntries[2].Weight)
	assert.NotNil(t, menuEntries[2].ID)

	/*
	 * Fourth menu entry
	 */
	assert.Equal(t, "plugin1", menuEntries[3].PluginName)
	assert.Equal(t, "menu 2", menuEntries[3].Name)
	assert.Equal(t, "route2", menuEntries[3].Route)
	assert.Equal(t, 5, menuEntries[3].Weight)
	assert.NotNil(t, menuEntries[3].ID)
}

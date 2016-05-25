package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/handlers"
	testUtils "github.com/eogile/agilestack-root-app/root-app-builder/server/testing"
	"github.com/eogile/agilestack-utils/plugins/components"
	"github.com/eogile/agilestack-utils/plugins/registration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	config1 = registration.PluginConfiguration{
		PluginName: "My wonderful plugin",
		Reducers: []string{
			"reducer1",
			"reducer2",
		},
		Routes: []registration.Route{
			registration.Route{
				ComponentName: "Component1",
				Href:          "/route-1",
				Routes:        []registration.SubRoute{},
				Type:          "content-route",
			},
			registration.Route{
				ComponentName: "Component1",
				Href:          "/route-2_1",
				Routes:        []registration.SubRoute{},
				Type:          "content-route",
			},
		},
	}

	config2 = registration.PluginConfiguration{
		PluginName: "Plugin 2",
		Reducers:   []string{},
		Routes: []registration.Route{
			registration.Route{
				ComponentName: "SomeBusinessComponent",
				Href:          "/route-10",
				Routes:        []registration.SubRoute{},
				Type:          "content-route",
			},
		},
	}
)

func TestHandlePluginsEndpoint_Routes(t *testing.T) {
	testUtils.DeleteAllStoreEntries(t)

	require.Nil(t, registration.StoreRoutesAndReducers(&config1))
	require.Nil(t, registration.StoreRoutesAndReducers(&config2))

	mux := http.NewServeMux()
	mux.HandleFunc("/", routesGenerationEndpoint())

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/", nil)

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

	var configurations []registration.PluginConfiguration
	err := json.Unmarshal(writer.Body.Bytes(), &configurations)
	require.Nil(t, err)
	require.Equal(t, 2, len(configurations))

	if configurations[0].PluginName == config1.PluginName {
		validateConfig(t, &config1, &configurations[0])
		validateConfig(t, &config2, &configurations[1])
	} else {
		validateConfig(t, &config2, &configurations[0])
		validateConfig(t, &config1, &configurations[1])
	}
}

func TestHandlePluginsEndpoint_NoRoutes(t *testing.T) {
	testUtils.DeleteAllStoreEntries(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/", routesGenerationEndpoint())

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/", nil)

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
	assert.Equal(t, "[]", string(writer.Body.Bytes()))
}

func TestHandlePluginsEndpoint_Components(t *testing.T) {
	testUtils.DeleteAllStoreEntries(t)

	storedComponent := &components.Components{
		PluginName:    "my-plugin",
		AppComponent:  "App3",
		MainComponent: "Main3",
	}
	require.Nil(t, components.StoreComponents(storedComponent))

	mux := http.NewServeMux()
	mux.HandleFunc("/", componentsGenerationEndpoint())

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/", nil)

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

	var result components.Components
	err := json.Unmarshal(writer.Body.Bytes(), &result)
	require.Nil(t, err)
	require.Equal(t, storedComponent.PluginName, result.PluginName)
	require.Equal(t, storedComponent.AppComponent, result.AppComponent)
	require.Equal(t, storedComponent.MainComponent, result.MainComponent)
}

func TestHandlePluginsEndpoint_NoComponents(t *testing.T) {
	testUtils.DeleteAllStoreEntries(t)

	mux := http.NewServeMux()
	mux.HandleFunc("/", componentsGenerationEndpoint())

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/", nil)

	mux.ServeHTTP(writer, request)

	/*
	 * Checking the HTTP status.
	 */
	assert.Equal(t, 404, writer.Code, "Invalid HTTP status")
}

func routesGenerationEndpoint() func(http.ResponseWriter, *http.Request) {
	return handlers.NewGenerationHandler(
		func(w http.ResponseWriter, configurations []registration.PluginConfiguration, AppComponents *components.Components) {
			bytes, err := json.Marshal(configurations)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(bytes)
			}
		}).HandlePluginsEndpoint
}

func componentsGenerationEndpoint() func(http.ResponseWriter, *http.Request) {
	return handlers.NewGenerationHandler(
		func(w http.ResponseWriter, configurations []registration.PluginConfiguration, appComponents *components.Components) {
			if appComponents == nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			bytes, err := json.Marshal(appComponents)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(bytes)
			}
		}).HandlePluginsEndpoint
}

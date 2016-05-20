package generators_test

import (
	"io/ioutil"
	"testing"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/generators"
	"github.com/eogile/agilestack-utils/plugins/registration"
	"github.com/ory-am/osin-storage/Godeps/_workspace/src/github.com/stretchr/testify/require"
)

/*
 * Tests the content of the generated file when there is no routes.
 *
 * In this case, an empty array is exported in JavaScript.
 */
func TestGenerateRoutesFileNoRoutes(t *testing.T) {
	fileName, fileClose := testTempFile(t)
	defer fileClose()

	err := generators.GenerateRoutesFile([]registration.PluginConfiguration{}, fileName)
	require.Nil(t, err, "Error while generating the file")

	bytes, err := ioutil.ReadFile(fileName)
	require.Nil(t, err, "Error while reading the generated file")

	expected := "Object.defineProperty(exports, \"__esModule\", {\n"
	expected += "  value: true\n"
	expected += "});\n"
	expected += "exports.default = [];\n"

	require.Equal(t, expected, string(bytes),
		"The file content does not match")
}

/*
 * Tests the content of the generated file when there is several routes.
 */
func TestGenerateRoutesFile(t *testing.T) {
	fileName, fileClose := testTempFile(t)
	defer fileClose()
	configurations := []registration.PluginConfiguration{
		registration.PluginConfiguration{
			PluginName: "module1",
			Reducers:   []string{},
			Routes:     []registration.Route{
				registration.Route{
					Href:          "route10",
					ComponentName: "component1",
				},
				registration.Route{
					Href:          "route20",
					ComponentName: "component2",
				},
			},
		},
		registration.PluginConfiguration{
			PluginName: "module2",
			Reducers:   []string{},
			Routes:     []registration.Route{
				registration.Route{
					Href:          "route30",
					ComponentName: "Component3",
				},
			},
		},
	}

	err := generators.GenerateRoutesFile(configurations, fileName)
	require.Nil(t, err, "Error while generating the file")

	bytes, err := ioutil.ReadFile(fileName)
	require.Nil(t, err, "Error while reading the generated file")

	/*
	 * The expected file content.
	 */
	expected := "Object.defineProperty(exports, \"__esModule\", {\n"
	expected += "  value: true\n"
	expected += "});\n"
	expected += "var _routeComponent00 = require('module1').component1;\n"
	expected += "var _route00 = {href: 'route10', component: _routeComponent00};\n"
	expected += "var _routeComponent01 = require('module1').component2;\n"
	expected += "var _route01 = {href: 'route20', component: _routeComponent01};\n"
	expected += "var _routeComponent10 = require('module2').Component3;\n"
	expected += "var _route10 = {href: 'route30', component: _routeComponent10};\n"
	expected += "exports.default = [_route00, _route01, _route10];\n"

	require.Equal(t, expected, string(bytes),
		"The file content does not match")
}

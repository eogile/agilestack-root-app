package generators_test

import (
	"io/ioutil"
	"testing"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/generators"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/models"
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

	err := generators.GenerateRoutesFile([]models.Route{}, fileName)
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

	routes := []models.Route{
		{
			Href:          "route10",
			ComponentName: "component1",
			ModuleName:    "module1",
		}, {
			Href:          "route30",
			ComponentName: "Component3",
			ModuleName:    "module2",
		},
		{
			Href:          "route20",
			ComponentName: "component2",
			ModuleName:    "module1",
		},
	}

	err := generators.GenerateRoutesFile(routes, fileName)
	require.Nil(t, err, "Error while generating the file")

	bytes, err := ioutil.ReadFile(fileName)
	require.Nil(t, err, "Error while reading the generated file")

	/*
	 * The expected file content.
	 */
	expected := "Object.defineProperty(exports, \"__esModule\", {\n"
	expected += "  value: true\n"
	expected += "});\n"
	expected += "var _routeComponent0 = require('module1').component1;\n"
	expected += "var _route0 = {href: 'route10', component: _routeComponent0};\n"
	expected += "var _routeComponent1 = require('module2').Component3;\n"
	expected += "var _route1 = {href: 'route30', component: _routeComponent1};\n"
	expected += "var _routeComponent2 = require('module1').component2;\n"
	expected += "var _route2 = {href: 'route20', component: _routeComponent2};\n"
	expected += "exports.default = [_route0, _route1, _route2];\n"


	require.Equal(t, expected, string(bytes),
		"The file content does not match")
}

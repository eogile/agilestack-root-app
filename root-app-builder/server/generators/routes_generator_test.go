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
			Routes: []registration.Route{
				registration.Route{
					Href:          "route10",
					ComponentName: "component1",
					Type:          "content-route",
					IsIndex:       true,
				},
				registration.Route{
					Href:          "route20",
					ComponentName: "component2",
					Type:          "full-screen-route",
					IsIndex:       false,
				},
			},
		},
		registration.PluginConfiguration{
			PluginName: "module2",
			Reducers:   []string{},
			Routes: []registration.Route{
				registration.Route{
					Href:          "route30",
					ComponentName: "Component3",
					Type:          "content-route",
					IsIndex:       false,
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
	expected += "exports.default = [\n"
	expected += "{href:'route10', isIndex:true, type:'content-route', component: require('module1').component1, routes:[]},\n"
	expected += "{href:'route20', isIndex:false, type:'full-screen-route', component: require('module1').component2, routes:[]},\n"
	expected += "{href:'route30', isIndex:false, type:'content-route', component: require('module2').Component3, routes:[]}"
	expected += "];\n"

	require.Equal(t, expected, string(bytes),
		"The file content does not match")
}

// With sub routes
func TestGenerateRoutesFile_WithSubRoutes(t *testing.T) {
	fileName, fileClose := testTempFile(t)
	defer fileClose()
	configurations := []registration.PluginConfiguration{
		registration.PluginConfiguration{
			PluginName: "module1",
			Reducers:   []string{},
			Routes: []registration.Route{
				registration.Route{
					Href:          "route10",
					ComponentName: "component1",
					Routes:        []registration.SubRoute{},
					Type:          "content-route",
					IsIndex:       false,
				},
				registration.Route{
					Href:          "route20",
					ComponentName: "component2",
					Routes: []registration.SubRoute{
						registration.SubRoute{
							Href:          "route201",
							ComponentName: "component21",
						},
						registration.SubRoute{
							Href:          "route202",
							ComponentName: "component22",
						},
					},
					Type: "full-screen-route",
				},
			},
		},
		registration.PluginConfiguration{
			PluginName: "module2",
			Reducers:   []string{},
			Routes: []registration.Route{
				registration.Route{
					ComponentName: "Component3",
					Routes: []registration.SubRoute{
						registration.SubRoute{
							Href:          "route301",
							ComponentName: "component31",
							Routes:        []registration.SubRoute{},
						},
					},
					Type:    "content-route",
					IsIndex: true,
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
	expected += "exports.default = [\n"
	expected += "{href:'route10', isIndex:false, type:'content-route', component: require('module1').component1, routes:[]},\n"
	expected += "{href:'route20', isIndex:false, type:'full-screen-route', component: require('module1').component2, routes:[{href:'route201', component: require('module1').component21, routes:[]}, {href:'route202', component: require('module1').component22, routes:[]}]},\n"
	expected += "{isIndex:true, type:'content-route', component: require('module2').Component3, routes:[{href:'route301', component: require('module2').component31, routes:[]}]}"
	expected += "];\n"

	require.Equal(t, expected, string(bytes),
		"The file content does not match")
}

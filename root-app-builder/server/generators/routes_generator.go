package generators

import (
	"strconv"

	"github.com/eogile/agilestack-utils/plugins/registration"
)

/*
 * Generating the JavaScript file that will contain the routes.
 */
func GenerateRoutesFile(configs []registration.PluginConfiguration, fileName string) error {
	return generateFile(generateRoutes(configs), fileName)
}

func generateRoutes(configs []registration.PluginConfiguration) string {
	result := "Object.defineProperty(exports, \"__esModule\", {\n"
	result += "  value: true\n"
	result += "});\n"

	if len(configs) == 0 {
		return result + "exports.default = [];\n"
	}

	/*
	 * Route JSON objects
	 */
	routesNames := ""
	for index, config := range configs {
		for index2, route := range config.Routes {
			result += "var _routeComponent" +
				strconv.Itoa(index) +
				strconv.Itoa(index2) +
				" = require('" +
				config.PluginName +
				"')." +
				route.ComponentName +
				";\n"
			result += "var _route" +
				strconv.Itoa(index) +
				strconv.Itoa(index2) +
				" = {href: '" +
				route.Href +
				"', component: _routeComponent" +
				strconv.Itoa(index) +
				strconv.Itoa(index2) +
				"};\n"

			routeName := "_route" + strconv.Itoa(index) + strconv.Itoa(index2)
			if index > 0 || index2 > 0{
				routesNames += ", "
			}
			routesNames += routeName
		}
	}

	return result + "exports.default = [" + routesNames + "];\n"
}

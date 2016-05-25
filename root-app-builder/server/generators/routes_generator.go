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

	result += "exports.default = [\n"

	for index1, config := range configs {
		for index2, route := range config.Routes {
			if index1 > 0 || index2 > 0 {
				result += ",\n"
			}

			result += routeJSObject(config.PluginName, route)
		}
	}

	return result + "];\n"
}

func routeJSObject(pluginName string, route registration.Route) string {
	result := "{"
	if route.Href != "" {
		result += "href:'" + route.Href + "', "
	}
	result += "isIndex:" + strconv.FormatBool(route.IsIndex) + ", "
	result += "type:'" + route.Type + "', "
	result += "component: require('" + pluginName + "')." + route.ComponentName + ", "
	result += "routes:["
	for index, subRoute := range route.Routes {
		if index != 0 {
			result += ", "
		}
		result += subRouteJSObject(pluginName, subRoute)
	}
	result += "]"
	return result + "}"
}

func subRouteJSObject(pluginName string, subRoute registration.SubRoute) string {
	result := "{"

	if subRoute.Href != "" {
		result += "href:'" + subRoute.Href + "', "
	}
	result += "component: require('" + pluginName + "')." + subRoute.ComponentName + ", "
	result += "routes:["
	for index, subSubRoute := range subRoute.Routes {
		if index != 0 {
			result += ", "
		}
		result += subRouteJSObject(pluginName, subSubRoute)
	}

	result += "]"
	return result + "}"
}

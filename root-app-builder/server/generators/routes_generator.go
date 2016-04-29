package generators

import (
	"strconv"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/models"
)

/*
 * Generating the JavaScript file that will contain the routes.
 */
func GenerateRoutesFile(routes []models.Route, fileName string) error {
	return generateFile(generateRoutes(routes), fileName)
}

func generateRoutes(routes []models.Route) string {
	result := "Object.defineProperty(exports, \"__esModule\", {\n"
	result += "  value: true\n"
	result += "});\n"

	if len(routes) == 0 {
		return result + "exports.default = [];\n"
	}

	/*
	 * Route JSON objects
	 */
	routesNames := ""
	for index, route := range routes {
		result += "var _routeComponent" +
			strconv.Itoa(index) +
			" = require('" +
			route.ModuleName +
			"')." +
			route.ComponentName +
			";\n"
		result += "var _route" +
			strconv.Itoa(index) +
			" = {href: '" +
			route.Href +
			"', component: _routeComponent" +
			strconv.Itoa(index) +
			"};\n"

		routeName := "_route" + strconv.Itoa(index)
		if index > 0 {
			routesNames += ", "
		}
		routesNames += routeName
	}

	return result + "exports.default = [" + routesNames + "];\n"
}
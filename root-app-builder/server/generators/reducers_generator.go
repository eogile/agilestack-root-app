package generators

import (
	"strconv"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/models"
)

/*
 * Generating the JavaScript file that will contain the reducers.
 */
func GenerateReducersFile(reducers []models.Reducer, fileName string) error {
	return generateFile(generateReducers(reducers), fileName)
}

func generateReducers(reducers []models.Reducer) string {
	result := "Object.defineProperty(exports, \"__esModule\", {\n"
	result += "  value: true\n"
	result += "});\n"

	if len(reducers) == 0 {
		return result + "exports.default = {};\n"
	}
	reducersNames := ""

	/*
	 * Generates the "import" statements for each reducer.
	 */
	for index, reducer := range reducers {
		result += "var _reducer" +
			strconv.Itoa(index) +
			" = require('" +
			reducer.ModuleName +
			"')." +
			reducer.Name +
			";\n"
		//result += generateImportStatement(reducer.Name, reducer.ModuleName)

		if index > 0 {
			reducersNames += ", "
		}
		reducersNames += reducer.Name + ": " + "_reducer" + strconv.Itoa(index)
	}

	return result + "exports.default = {" + reducersNames + "};\n"
}

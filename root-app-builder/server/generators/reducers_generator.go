package generators

import (
	"strconv"

	"github.com/eogile/agilestack-utils/plugins/registration"
)

/*
 * Generating the JavaScript file that will contain the reducers.
 */
func GenerateReducersFile(configs []registration.PluginConfiguration, fileName string) error {
	return generateFile(generateReducers(configs), fileName)
}

func generateReducers(configs []registration.PluginConfiguration) string {
	result := "Object.defineProperty(exports, \"__esModule\", {\n"
	result += "  value: true\n"
	result += "});\n"

	if len(configs) == 0 {
		return result + "exports.default = {};\n"
	}
	reducersNames := ""

	/*
	 * Generates the "import" statements for each reducer.
	 */
	for index, config := range configs {
		for index2, reducer := range config.Reducers {
			result += "var _reducer" +
				strconv.Itoa(index) +
				strconv.Itoa(index2) +
				" = require('" +
				config.PluginName +
				"')." +
				reducer +
				";\n"

			if index > 0 || index2 > 0 {
				reducersNames += ", "
			}
			reducersNames += reducer + ": " + "_reducer" + strconv.Itoa(index) + strconv.Itoa(index2)
		}
	}

	return result + "exports.default = {" + reducersNames + "};\n"
}

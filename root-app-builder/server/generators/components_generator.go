package generators

import (
	"github.com/eogile/agilestack-utils/plugins/components"
)

func GenerateComponentsFile(components *components.Components, fileName string) error {
	return generateFile(generateComponents(components), fileName)
}

func generateComponents(components *components.Components) string {
	result := "Object.defineProperty(exports, \"__esModule\", {\n"
	result += "  value: true\n"
	result += "});\n"

	if components == nil {
		return result + defaultComponents()
	}

	result += "exports.App = require('" + components.PluginName + "')." + components.AppComponent + ";\n";
	result += "exports.Main = require('" + components.PluginName + "')." + components.MainComponent + ";\n";

	return result
}

func defaultComponents() string {
	result := "var App = require('../components/App.react').default;\n"
	result += "var Main = require('../components/Main.react').default;\n"
	result+= "exports.App = App;\n"
	result+= "exports.Main = Main;\n"
	return result
}

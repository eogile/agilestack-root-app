package generators_test

import (
	"io/ioutil"
	"testing"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/generators"
	"github.com/eogile/agilestack-utils/plugins/components"
	"github.com/stretchr/testify/require"
)

func TestGenerateComponentsFile_NoComponents(t *testing.T) {
	fileName, closeFunc := testTempFile(t)
	defer closeFunc()

	require.Nil(t, generators.GenerateComponentsFile(nil, fileName))

	bytes, err := ioutil.ReadFile(fileName)
	require.Nil(t, err)

	expected := "Object.defineProperty(exports, \"__esModule\", {\n"
	expected += "  value: true\n"
	expected += "});\n"
	expected += "var component = require('../components/App.react').default;\n"
	expected += "var component = require('../components/Main.react').default;\n"
	expected += "exports.App = component;\n"
	expected += "exports.Main = component;\n"
	require.Equal(t, expected, string(bytes))
}

func TestGenerateComponentsFile(t *testing.T) {
	fileName, closeFunc := testTempFile(t)
	defer closeFunc()

	require.Nil(t, generators.GenerateComponentsFile(&components.Components{
		PluginName:    "my-plugin",
		AppComponent:  "App2",
		MainComponent: "Main2",
	}, fileName))

	bytes, err := ioutil.ReadFile(fileName)
	require.Nil(t, err)

	expected := "Object.defineProperty(exports, \"__esModule\", {\n"
	expected += "  value: true\n"
	expected += "});\n"
	expected += "exports.App = require('my-plugin').App2;\n"
	expected += "exports.Main = require('my-plugin').Main2;\n"
	require.Equal(t, expected, string(bytes))
}

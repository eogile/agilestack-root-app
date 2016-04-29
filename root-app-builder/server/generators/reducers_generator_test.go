package generators_test

import (
	"io/ioutil"
	"testing"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/generators"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/models"
	"github.com/stretchr/testify/require"
)

/*
 * Tests the content of the generated file when there is no reducers.
 *
 * In this case, an empty array is exported in JavaScript.
 */
func TestGenerateReducersFileNoReducer(t *testing.T) {
	fileName, closeFunc := testTempFile(t)
	defer closeFunc()

	err := generators.GenerateReducersFile([]models.Reducer{}, fileName)
	require.Nil(t, err)

	bytes, err := ioutil.ReadFile(fileName)
	require.Nil(t, err)

	expected := "Object.defineProperty(exports, \"__esModule\", {\n";
	expected += "  value: true\n"
	expected += "});\n"
	expected += "exports.default = {};\n"
	require.Equal(t, expected, string(bytes))
}

/*
 * Tests the content of the generated file when there is
 * several reducers.
 */
func TestGenerateReducersFile(t *testing.T) {
	fileName, closeFunc := testTempFile(t)
	defer closeFunc()

	reducers := []models.Reducer{
		{
			Name: "Reducer1",
			ModuleName: "module2",
		},
		{
			Name: "reducer2",
			ModuleName: "module2",
		},
		{
			Name: "ReDuCeR3",
			ModuleName: "module3",
		},
	}

	expected := "Object.defineProperty(exports, \"__esModule\", {\n";
	expected += "  value: true\n"
	expected += "});\n"
	expected += "var _reducer0 = require('module2').Reducer1;\n"
	expected += "var _reducer1 = require('module2').reducer2;\n"
	expected += "var _reducer2 = require('module3').ReDuCeR3;\n"
	expected += "exports.default = {Reducer1: _reducer0, reducer2: _reducer1, ReDuCeR3: _reducer2};\n"

	err := generators.GenerateReducersFile(reducers, fileName)
	require.Nil(t, err)

	bytes, err := ioutil.ReadFile(fileName)
	require.Nil(t, err)
	require.Equal(t, expected, string(bytes))
}

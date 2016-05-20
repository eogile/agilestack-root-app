package handlers_test

import (
	"testing"

	testUtils "github.com/eogile/agilestack-root-app/root-app-builder/server/testing"
	"github.com/eogile/agilestack-utils/plugins/registration"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	testUtils.Main(m)
}

func validateConfig(t *testing.T, expectedConfig, resultConfig *registration.PluginConfiguration) {
	require.Equal(t, expectedConfig.PluginName, resultConfig.PluginName)
	require.Equal(t, expectedConfig.Reducers, resultConfig.Reducers)
	require.Equal(t, len(expectedConfig.Routes), len(resultConfig.Routes))

	for i, expectedRoute := range expectedConfig.Routes {
		resultRoute := resultConfig.Routes[i]
		require.Equal(t, expectedRoute.ComponentName, resultRoute.ComponentName)
		require.Equal(t, expectedRoute.Href, resultRoute.Href)
	}
}

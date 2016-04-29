package repository_test

import (
	"testing"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	testUtils "github.com/eogile/agilestack-root-app/root-app-builder/server/testing"
)

/*
 * Tests that an empty slice is returned when there is no
 * menu entries.
 */
func TestFindAllNoMenuEntries(t *testing.T) {
	testUtils.SetUp(t)
	testUtils.TruncateTables(t)

	menuEntries, err := menuEntriesRepo.FindAll()
	require.Nil(t, err)
	require.Equal(t, 0, len(menuEntries))
}

/*
 * Tests that the menu entries are successfully loaded and
 * correctly sorted.
 */
func TestFindAllMenuEntries(t *testing.T) {
	testUtils.SetUp(t)

	menuEntries, err := menuEntriesRepo.FindAll()
	require.Nil(t, err)
	require.Equal(t, 3, len(menuEntries))

	assert.Equal(t, "plugin1", menuEntries[0].PluginName)
	assert.Equal(t, "menu 1", menuEntries[0].Name)
	assert.Equal(t, "route1", menuEntries[0].Route)
	assert.Equal(t, 10, menuEntries[0].Weight)
	assert.NotNil(t, menuEntries[0].ID)

	assert.Equal(t, "plugin2", menuEntries[1].PluginName)
	assert.Equal(t, "menu 3", menuEntries[1].Name)
	assert.Equal(t, "route3", menuEntries[1].Route)
	assert.Equal(t, 10, menuEntries[1].Weight)
	assert.NotNil(t, menuEntries[1].ID)

	assert.Equal(t, "plugin1", menuEntries[2].PluginName)
	assert.Equal(t, "menu 2", menuEntries[2].Name)
	assert.Equal(t, "route2", menuEntries[2].Route)
	assert.Equal(t, 5, menuEntries[2].Weight)
	assert.NotNil(t, menuEntries[2].ID)
}

/*
 * Tests the insertion of menu entries.
 *
 * There are already two menu entries for the given plugin :
 * 	- One is deleted,
 *	- One is updated.
 */
func TestSaveMenuEntries(t *testing.T) {
	testUtils.SetUp(t)

	entries := []models.MenuEntry{
		{
			Name:       "menu n1 new name",
			PluginName: "plugin1",
			Route:      "route1-new",
			Weight:     10,
		},
	}
	err := menuEntriesRepo.Save(entries)
	require.Nil(t, err)

	menuEntries, err := menuEntriesRepo.FindAll()
	require.Nil(t, err)
	require.Equal(t, 2, len(menuEntries))

	assert.Equal(t, "plugin2", menuEntries[0].PluginName)
	assert.Equal(t, "menu 3", menuEntries[0].Name)
	assert.Equal(t, "route3", menuEntries[0].Route)
	assert.Equal(t, 10, menuEntries[0].Weight)
	assert.NotNil(t, menuEntries[0].ID)

	assert.Equal(t, "menu n1 new name", menuEntries[1].Name)
	assert.Equal(t, "plugin1", menuEntries[1].PluginName)
	assert.Equal(t, "route1-new", menuEntries[1].Route)
	assert.Equal(t, 10, menuEntries[1].Weight)
	assert.NotNil(t, menuEntries[1].ID)
}

/*
 * Tests the rollback mechanism when an error is thrown in the
 * SQL transaction.
 *
 * The menu entries insertion has two steps :
 * 	1 - Delete existing menu entries,
 *	2 - Inserts the new menu entries.
 *
 * In this test, the first step works but a integrity constraint is
 * violated during the second step (routes must be unique).
 *
 * As a result, the stored menu entries must be exactly the same than
 * before the test.
 */
func TestSaveMenuEntriesError(t *testing.T) {
	testUtils.SetUp(t)

	/*
	 * The route is already registered for another plugin.
	 */
	entries := []models.MenuEntry{
		{
			Name:       "menu n1 new name",
			PluginName: "plugin1",
			Route:      "route3",
			Weight:     10,
		},
	}

	err := menuEntriesRepo.Save(entries)
	require.NotNil(t, err)


	menuEntries, err := menuEntriesRepo.FindAll()
	require.Nil(t, err)
	require.Equal(t, 3, len(menuEntries))

	/*
	 * The menu entries did not changed.
	 */
	assert.Equal(t, "menu 1", menuEntries[0].Name)
	assert.Equal(t, "route1", menuEntries[0].Route)
	assert.Equal(t, "menu 3", menuEntries[1].Name)
	assert.Equal(t, "route3", menuEntries[1].Route)
	assert.Equal(t, "menu 2", menuEntries[2].Name)
	assert.Equal(t, "route2", menuEntries[2].Route)
}

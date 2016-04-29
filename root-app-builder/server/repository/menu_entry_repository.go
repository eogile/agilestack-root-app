package repository

import (
	"github.com/eogile/agilestack-root-app/root-app-builder/server/models"
)

type MenuEntryRepository interface {

	/*
	 * Creates if needed the tables the repository is responsible for.
	 *
	 * If the table already exists, nothing must be done.
	 */
	CreateSchema() error

	/*
	 * Returns the list containing all the menu entries.
	 *
	 * The menu entries must be sorted by weight (descending) and
	 * name (ascending).
	 */
	FindAll() ([]models.MenuEntry, error)

	/*
	 * Saves the given menu entries into the database.
	 *
	 * All the given menu entries must be related to an unique
	 * plugin.
	 *
	 * All the previous menu entries related to the plugin are
	 * removed.
	 */
	Save([]models.MenuEntry) error
}

var repo = &PostgresMenuEntryRepository{}

/*
 * Returns the repository for menu entries to use.
 */
func GetRepository() MenuEntryRepository {
	return repo
}

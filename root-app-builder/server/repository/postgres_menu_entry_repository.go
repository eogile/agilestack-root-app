package repository

import (
	"database/sql"
	"github.com/pborman/uuid"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/models"
	"log"
)

type PostgresMenuEntryRepository struct {
}

func (repo *PostgresMenuEntryRepository) CreateSchema() error {
	menuEntriesSchema := `CREATE TABLE IF NOT EXISTS MENU_ENTRIES (
		id          text NOT NULL PRIMARY KEY,
		plugin_name varchar(64) NOT NULL,
		name        varchar(64) NOT NULL UNIQUE,
		route       varchar(64) NOT NULL UNIQUE,
		weight	    int
	)`

	_, err := db.Exec(menuEntriesSchema)
	if  err != nil {
		log.Println("Error while creating the MENU_ENTRIES schema")
	}
	return err
}


func (repo *PostgresMenuEntryRepository) FindAll() ([]models.MenuEntry, error) {
	selectQuery := `
		SELECT id, name, plugin_name, route, weight
			FROM MENU_ENTRIES
			ORDER BY weight DESC, name ASC
	`

	rows, err := db.Query(selectQuery)
	if err != nil {
		return nil, err
	}

	menuEntries := make([]models.MenuEntry, 0)
	for rows.Next() {
		menuEntry := new(models.MenuEntry)
		err = rows.Scan(
			&menuEntry.ID,
			&menuEntry.Name,
			&menuEntry.PluginName,
			&menuEntry.Route,
			&menuEntry.Weight)

		if err != nil {
			return nil, err
		}
		menuEntries = append(menuEntries, *menuEntry)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return menuEntries, nil
}

func (repo *PostgresMenuEntryRepository) Save(menuEntries []models.MenuEntry) error {
	insertRequest := `
	   insert into MENU_ENTRIES
	   	(id, plugin_name, name, route, weight)
	   	values ($1, $2, $3, $4, $5)
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("delete from MENU_ENTRIES where plugin_name = $1",
		menuEntries[0].PluginName)
	if err != nil {
		return rollback(tx, err)
	}

	for _, menuEntry := range menuEntries {
		menuEntry.ID = uuid.New()
		_, err = tx.Exec(insertRequest,
			menuEntry.ID,
			menuEntry.PluginName,
			menuEntry.Name,
			menuEntry.Route,
			menuEntry.Weight)
		if err != nil {
			return rollback(tx, err)
		}
	}

	return tx.Commit()
}

func rollback(tx *sql.Tx, err error) error {
	log.Println("Error in SQL transaction: ", err)
	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		log.Println("Error while rollbacking: ", rollbackErr)
	}
	return err
}

package testing


import (
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/eogile/agilestack-root-app/root-app-builder/server/repository"
	"gopkg.in/ory-am/dockertest.v2"
	"os"
	"github.com/ory-am/osin-storage/Godeps/_workspace/src/github.com/stretchr/testify/require"
)

var menuEntriesRepo *repository.PostgresMenuEntryRepository
var testDB *sql.DB

func Main(m *testing.M) {
	/*
	 * Bootstrap the Postgres container.
	 */
	db, removeContainerFunc := startPostgres()
	testDB = db
	repository.SetDB(testDB)

	/*
	 * Initializing the repositories
	 */
	menuEntriesRepo = &repository.PostgresMenuEntryRepository{}
	err := menuEntriesRepo.CreateSchema()
	if err != nil {
		log.Fatalln("Error while creating the schema: ", err)
	}

	exitCode := m.Run()

	/*
	 * Explicit call to the "clean" function. Indeed, it can't used
	 * with "defer removeContainerFunc()" because it won't be executed
	 * due to the "os.Exit()" call.
	 */
	removeContainerFunc()
	os.Exit(exitCode)
}

func startPostgres() (*sql.DB, func()) {
	var dbObject *sql.DB

	c, err := dockertest.ConnectToPostgreSQL(15, time.Second, func(url string) bool {
		var err error
		dbObject, err = sql.Open("postgres", url)
		if err != nil {
			return false
		}

		return dbObject.Ping() == nil
	})

	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
		return nil, nil
	}

	return dbObject, func() {
		dbObject.Close()
		c.KillRemove()
	}
}

func SetUp(t *testing.T) {
	TruncateTables(t)

	/*
	 * Inserting data
	 */
	_, err := testDB.Exec(`INSERT INTO MENU_ENTRIES (id, plugin_name, name, route, weight)
		VALUES
		('1', 'plugin1', 'menu 1', 'route1', 10),
		('2', 'plugin1', 'menu 2', 'route2', 5),
		('3', 'plugin2', 'menu 3', 'route3', 10);
	`);
	require.Nil(t, err)
}

func TruncateTables(t *testing.T) {
	/*
	 * Removing all existing data
	 */
	_, err := testDB.Exec("DELETE FROM MENU_ENTRIES")
	require.Nil(t, err)
}


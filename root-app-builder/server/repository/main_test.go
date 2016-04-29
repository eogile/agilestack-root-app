package repository_test

import (
	"testing"
	"github.com/eogile/agilestack-root-app/root-app-builder/server/repository"
	testUtils "github.com/eogile/agilestack-root-app/root-app-builder/server/testing"
)

var menuEntriesRepo *repository.PostgresMenuEntryRepository

func TestMain(m *testing.M) {
	menuEntriesRepo = &repository.PostgresMenuEntryRepository{}
	testUtils.Main(m)
}

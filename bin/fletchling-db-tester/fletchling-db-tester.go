package main

import (
	"context"
	"log"
	"os"

	"github.com/UnownHash/Fletchling/app_config"
	"github.com/UnownHash/Fletchling/db_store"
	"github.com/UnownHash/Fletchling/version"
)

const (
	LOGFILE_NAME                  = "fletchling-db-refresher.log"
	DEFAULT_CONFIG_FILENAME       = "./configs/fletchling.toml"
	DEFAULT_NESTS_MIGRATIONS_PATH = "./db_store/sql"
)

func main() {
	defaultConfig := app_config.GetDefaultConfig()

	cfg, err := app_config.LoadConfig(DEFAULT_CONFIG_FILENAME, defaultConfig)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Logging.Filename = LOGFILE_NAME

	logger := cfg.CreateLogger(true)
	logger.Infof("STARTUP: Version %s. Config loaded.", version.APP_VERSION)

	// check destination first before we attempt to load
	// area fences.
	nestsDBStore, err := db_store.NewNestsDBStore(cfg.NestsDb, logger)
	if err != nil {
		logger.Errorf("failed to init nests db for db importer: %v", err)
		os.Exit(1)
	}

	if _, _, err := nestsDBStore.CheckMigrate(DEFAULT_NESTS_MIGRATIONS_PATH); err != nil {
		logger.Errorf("error initing nests db: %v", err)
		os.Exit(1)
	}

	err = nestsDBStore.DebugPolygons(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
}

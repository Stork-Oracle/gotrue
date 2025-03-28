package test

import (
	"github.com/stork-oracle/gotrue/internal/conf"
	"github.com/stork-oracle/gotrue/internal/storage"
)

func SetupDBConnection(globalConfig *conf.GlobalConfiguration) (*storage.Connection, error) {
	return storage.Dial(globalConfig)
}

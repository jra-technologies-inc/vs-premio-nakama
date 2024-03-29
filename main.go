package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	initStart := time.Now()

	err := initializer.RegisterRpc("healthcheck", RpcHealthcheck)

	logger.Info("Module loaded in %dms", time.Since(initStart).Milliseconds())
	return err
}

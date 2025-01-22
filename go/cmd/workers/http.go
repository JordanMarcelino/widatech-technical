package workers

import (
	"context"

	"github.com/JordanMarcelino/widatech-technical/internal/config"
	"github.com/JordanMarcelino/widatech-technical/internal/server"
)

func runHttpWorker(cfg *config.Config, ctx context.Context) {
	srv := server.NewHttpServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}

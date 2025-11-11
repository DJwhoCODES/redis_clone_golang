package main

import (
	"github.com/djwhocodes/redis-clone/internal/logger"
	"github.com/djwhocodes/redis-clone/internal/server"
	"github.com/djwhocodes/redis-clone/internal/store"
)

func main() {
	logger.Init()

	store := store.NewStore()
	handler := server.NewHandler(store)
	srv := server.NewServer(":5000", handler)

	if err := srv.Start(); err != nil {
		logger.Error.Fatal(err)
	}
}

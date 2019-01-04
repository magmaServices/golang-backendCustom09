package main

// Generate setters and getters for stats
// $ go generate ./cmd/fesl-backend
//go:generate go run ../stats-codegen/main.go -scan ../../backend/model --getters ../../backend/ranking/getters.go --setters ../../backend/ranking/setters.go --adders ../../backend/ranking/adders.go

import (
	"context"
	"flag"

	"github.com/google/gops/agent"
	"github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"

	"localhost/go-heroes/fesl-backend/backend/config"
	"localhost/go-heroes/fesl-backend/backend/fesl"
	"localhost/go-heroes/fesl-backend/backend/matchmaking"
	"localhost/go-heroes/fesl-backend/backend/network"
	"localhost/go-heroes/fesl-backend/backend/storage/database"
	"localhost/go-heroes/fesl-backend/backend/storage/kvstore"
	"localhost/go-heroes/fesl-backend/backend/theater"
)

func main() {
	initConfig()
	initLogger()

	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()
	startServer(ctx)

	// Use "github.com/google/gops/agent"
	if err := agent.Listen(agent.Options{}); err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("Serving...")
	<-ctx.Done()
}

func initConfig() {
	var (
		configFile string
	)
	flag.StringVar(&configFile, "config", ".env", "Path to configuration file")
	flag.Parse()

	gotenv.Load(configFile)
	config.Initialize()
}

func initLogger() {
	logrus.SetLevel(config.LogLevel())

	// logrus.SetFormatter(&logrus.JSONFormatter{
	// 	DisableTimestamp: true,
	// })
	// logrus.SetFormatter(new(prefixed.TextFormatter))
	// logrus.SetFormatter(&prefixed.TextFormatter{
	// 	DisableTimestamp: true,
	// 	DisableColors:    true,
	// })
}

func startServer(ctx context.Context) {
	db, err := database.New()
	if err != nil {
		logrus.Fatal(err)
	}

	network.InitClientData()
	kvs := kvstore.NewInMemory()
	mm := matchmaking.NewPool()

	fesl.New(config.FeslClientAddr(), false, db, kvs, mm).ListenAndServe(ctx)
	fesl.New(config.FeslServerAddr(), true, db, kvs, mm).ListenAndServe(ctx)

	theater.New(config.ThtrClientAddr(), db, kvs, mm).ListenAndServe(ctx)
	theater.New(config.ThtrServerAddr(), db, kvs, mm).ListenAndServe(ctx)
}

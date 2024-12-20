package main

import (
	"flag"
	"log"

	"web-10/internal/query/api"
	"web-10/internal/query/config"
	"web-10/internal/query/provider"
	"web-10/internal/query/usecase"

	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config-path", "/home/balu/Рабочий стол/web-10-master/configs/query_example.yaml", "Path to configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	uc := usecase.NewUsecase(prv)
	srv := api.NewServer(cfg.IP, cfg.Port, uc)
	srv.Run()
}

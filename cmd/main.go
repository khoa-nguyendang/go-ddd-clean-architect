package main

import (
	config "app/core/configs"
	"app/pkg/logger"
	"app/pkg/mysql"
	"flag"
	"log"
	"os"
	"strconv"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
}

func main() {
	var (
		port           = flag.Int("port", 8080, "The service port")
		configfilepath = flag.String("configfilepath", "", "The service config")
	)
	flag.Parse()
	var configPath string
	if configfilepath != nil && *configfilepath != "" {
		configPath = *configfilepath
	} else {
		//load default
		configPath = config.GetConfigPath(os.Getenv("config"))
		log.Printf("Loadding default config path : %v \n", configPath)
	}
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config err: %v", err)
		return
	}

	var new_port = os.Args[1]
	logger := logger.NewLogger(logger.LoggerConfig(cfg.Logger))
	logger.InitLogger()
	srv, err := NewApp(cfg, logger)

	if err != nil {
		log.Printf("unable to resolve new application: %v", err)
		return
	}

	err = mysql.RunMigration(cfg)
	if err != nil {
		logger.Errorf("MySql RunMigration: %s", err)
	}

	i, err := strconv.Atoi(new_port)
	if err != nil {
		if err = srv.Run(*port); err != nil {
			log.Printf("run service error: %v", err)
		}
	} else {
		if err = srv.Run(i); err != nil {
			log.Printf("run service error: %v", err)
		}
	}

}

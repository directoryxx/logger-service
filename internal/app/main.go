package main

import (
	"log"
	"logger/config"
	"logger/delivery/http"
	"logger/delivery/worker"
	"os"
)

func main() {

	envSource := "SYSTEM"

	if os.Getenv("BYPASS_ENV_FILE") == "" {
		log.Println("[INFO] Load Config")
		config.LoadConfig()
		envSource = "FILE"
	}

	log.Println("[INFO] Loaded Config : " + envSource)

	if os.Getenv("APPLICATION_MODE") == "worker" {
		worker.RunWorkerServiceLog()
	} else {
		http.RunAPI()
	}
}

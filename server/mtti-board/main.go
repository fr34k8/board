package main

import (
	"os"
	"log"
	"github.com/mtti/board/server/daemon"
	"github.com/mtti/board/server/repositories"
)

func main() {

	// Create server
	s, err := daemon.New()
	if err != nil {
		log.Print("Server intialization failed: " + err.Error())
		os.Exit(1)
	}

	// Create default MySQL repository
	repo, err := repositories.NewMySQL(s.Settings.DatabaseURL)
	if err != nil {
		log.Print("Repository creation failed: " + err.Error())
		os.Exit(1)
	}
	s.Context.Repository = repo

	err = s.Start()
	if err != nil {
		s.Context.Log.Fatal(err)
		os.Exit(1)
	}

}

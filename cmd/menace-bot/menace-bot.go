package main

import (
	"flag"
	"os"

	"github.com/Plloi/pdb-cmdr/pkg/router"
	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

// Variables used for command line parameters
var (
	Token  string
	Router *router.CommandRouter
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Discord Token")
	flag.Parse()
	log.SetLevel(log.DebugLevel)
	godotenv.Load()

	if Token == "" {
		Token = os.Getenv("DISCORD_TOKEN")
	}

}

func main() {
	// TODO Basic Bot to play Menace
}

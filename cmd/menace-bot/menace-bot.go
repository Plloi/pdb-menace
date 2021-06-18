package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/Plloi/pdb-cmdr/pkg/router"
	"github.com/Plloi/pdb-menace/pkg/menace"
	"github.com/bwmarrin/discordgo"
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
	log.Info("Starting up...")
	// Create a new Discord session using the provided bot token.
	if Token == "" {
		log.Fatalln("ERROR: Token Required for bot usage")
	}
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatalf("Error creating Discord session: %v\n", err)
	}

	Router = router.NewCommandRouter()
	Router.DefaultPrefix = "m!"
	Router.RegisterCommand("prefix", "Sets the bot command prefix (Admin Locked)", Router.SetPrefix)

	// Register router's command handler for message events.
	dg.AddHandler(Router.HandleCommand)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection: %v\n", err)
	}

	// Cleanly close down the Discord session at end of exectution.
	defer dg.Close()

	// Load Menace Commands
	menace.Setup(Router)

	// Wait here until CTRL-C or other term signal is received.
	log.Info("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
